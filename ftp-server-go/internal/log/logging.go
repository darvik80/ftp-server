package log

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"ftp-server-go/internal/config"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ANSI-цвета, как в logback AnsiColorConverter.
const (
	reset   = "\x1b[0m"
	faint   = "\x1b[2m"
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
)

// springTime = %d{yyyy-MM-dd'T'HH:mm:ss.SSSXXX}
const springTime = "2006-01-02T15:04:05.000-07:00"

type ctxKey struct{}

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*slog.Logger); ok && l != nil {
		return l
	}

	return slog.Default()
}

func FromContextNamed(ctx context.Context, name string) *slog.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*slog.Logger); ok && l != nil {
		return l.With("logger", name)
	}

	return slog.Default()
}

func Setup(cfg config.LoggerConfig) (*slog.Logger, func() error, error) {
	root, err := parseLevel(cfg.Level)
	if err != nil {
		return nil, nil, err
	}

	var (
		handlers []slog.Handler
		flushers []*bufio.Writer
		closers  []io.Closer
	)

	consoleLevel, err := atLeast(root, cfg.ConsoleLevel)
	if err != nil {
		return nil, nil, err
	}
	switch strings.ToLower(cfg.Console) {
	case "spring":
		handlers = append(handlers, &springHandler{
			w:         os.Stdout,
			mu:        &sync.Mutex{},
			level:     consoleLevel,
			color:     useColor(cfg.Color),
			addSource: cfg.AddSource,
			pid:       os.Getpid(),
			logger:    "app",
			thread:    "main",
		})
	case "json":
		handlers = append(handlers, slog.NewJSONHandler(os.Stdout, jsonOptions(consoleLevel, cfg.AddSource)))
	case "text":
		handlers = append(handlers, slog.NewTextHandler(os.Stdout, jsonOptions(consoleLevel, cfg.AddSource)))
	case "off":
	default:
		return nil, nil, fmt.Errorf("log: console format %q: want spring|json|text|off", cfg.Console)
	}

	// --- файл: JSON для логовой/ELK, уровень ниже ---
	if cfg.File != "" {
		fileLevel, err := atLeast(root, cfg.FileLevel)
		if err != nil {
			return nil, nil, err
		}
		rw, err := newRotate(cfg.File, cfg.MaxSizeMB, cfg.MaxBackups)
		if err != nil {
			return nil, nil, err
		}
		bw := bufio.NewWriterSize(rw, 64*1024)
		flushers = append(flushers, bw)
		closers = append(closers, rw)
		handlers = append(handlers, slog.NewJSONHandler(bw, jsonOptions(fileLevel, cfg.AddSource)))
	}

	if len(handlers) == 0 {
		handlers = append(handlers, slog.DiscardHandler)
	}

	logger := slog.New(slog.NewMultiHandler(handlers...))
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(max(root, slog.LevelInfo)) // log.Printf пишем как INFO

	close := func() error {
		var errs []error
		for _, w := range flushers {
			if err := w.Flush(); err != nil {
				errs = append(errs, fmt.Errorf("log flush: %w", err))
			}
		}
		for _, c := range closers {
			if err := c.Close(); err != nil {
				errs = append(errs, fmt.Errorf("log close: %w", err))
			}
		}
		return errors.Join(errs...)
	}

	return logger, close, nil
}

func jsonOptions(l slog.Level, addSource bool) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: addSource,
		Level:     l,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				if t, ok := a.Value.Any().(time.Time); ok && !t.IsZero() {
					return slog.String(slog.TimeKey, t.Format(springTime))
				}
			}
			return a
		},
	}
}

func parseLevel(s string) (slog.Level, error) {
	var l slog.Level
	if err := l.UnmarshalText([]byte(s)); err != nil {
		return 0, fmt.Errorf("log: level %q: %w", s, err)
	}
	return l, nil
}

// atLeast — уровень sink'а не может быть мягче общего минимума.
func atLeast(root slog.Level, s string) (slog.Level, error) {
	if s == "" {
		return root, nil
	}
	l, err := parseLevel(s)
	if err != nil {
		return 0, err
	}
	return max(l, root), nil
}

// useColor: NO_COLOR (де-факто стандарт) > явная настройка > TTY-детект.
func useColor(mode string) bool {
	switch strings.ToLower(mode) {
	case "always":
		return true
	case "never":
		return false
	}
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	fi, err := os.Stdout.Stat()
	return err == nil && fi.Mode()&os.ModeCharDevice != 0
}

type springHandler struct {
	w         io.Writer
	mu        *sync.Mutex // одна на все копии этого sink'а
	level     slog.Level
	color     bool
	addSource bool
	pid       int
	logger    string
	thread    string
	attrs     []slog.Attr
	groups    []string
}

func (h *springHandler) Enabled(_ context.Context, l slog.Level) bool { return l >= h.level }

func (h *springHandler) Handle(_ context.Context, r slog.Record) error {
	if r.Level < h.level {
		return nil
	}
	b := make([]byte, 0, 256+len(r.Message))

	b = r.Time.AppendFormat(b, springTime)
	b = append(b, ' ')
	b = h.paint(b, lvlColor(r.Level), fmt.Sprintf("%5s", lvlName(r.Level)))
	b = append(b, ' ')
	b = h.paint(b, magenta, strconv.Itoa(h.pid))
	b = append(b, ' ')
	b = h.paint(b, faint, "---")
	b = append(b, ' ')
	b = h.paint(b, faint, "["+padLeft(trunc(h.thread, 10), 10)+"]")
	b = append(b, ' ')
	b = h.paint(b, cyan, padRight(trunc(abbrev(h.logger, 10), 10), 10))
	b = append(b, ' ')
	b = h.paint(b, faint, ":")
	b = append(b, ' ')
	b = append(b, h.colorize(msgColor(r.Level), r.Message)...)

	for _, a := range h.attrs {
		b = h.appendAttr(b, a)
	}
	var stack string
	r.Attrs(func(a slog.Attr) bool {
		if s := stackOf(a.Value.Any()); s != "" {
			stack = s
			return true
		}
		b = h.appendAttr(b, a)
		return true
	})

	if h.addSource && r.Source != nil {
		src := r.Source()
		b = h.paint(b, faint, " ("+filepath.Base(src.File)+":"+strconv.Itoa(src.Line)+")")
	}
	if stack != "" {
		b = append(b, '\n')
		for _, line := range strings.Split(stack, "\n") {
			b = append(b, "    "...)
			b = append(b, line...)
			b = append(b, '\n')
		}
		b = bytesTrimSpace(b)
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(append(b, '\n'))
	return err
}

// WithAttrs возвращает копию: slog.Handler по контракту неизменяем.
// Ключи-колонки извлекаются в поля, остальное остаётся атрибутами.
func (h *springHandler) WithAttrs(as []slog.Attr) slog.Handler {
	n := *h
	n.attrs = append([]slog.Attr(nil), h.attrs...)
	for _, a := range as {
		switch a.Key {
		case "logger":
			n.logger = asString(a.Value)
		case "thread":
			n.thread = asString(a.Value)
		default:
			n.attrs = append(n.attrs, a)
		}
	}
	return &n
}

func (h *springHandler) WithGroup(name string) slog.Handler {
	n := *h
	n.groups = append(append([]string(nil), h.groups...), name)
	return &n
}

func (h *springHandler) appendAttr(b []byte, a slog.Attr) []byte {
	if a.Equal(slog.Attr{}) {
		return b
	}
	key := a.Key
	if len(h.groups) > 0 {
		key = strings.Join(append(append([]string(nil), h.groups...), a.Key), ".")
	}
	b = append(b, ' ')
	b = append(b, key...)
	b = append(b, '=')
	v := asString(a.Value)
	if strings.ContainsAny(v, " \t\"=") {
		v = strconv.Quote(v)
	}
	return append(b, v...)
}

func (h *springHandler) colorize(color, s string) string {
	if !h.color || color == "" {
		return s
	}
	return color + s + reset
}

func (h *springHandler) paint(b []byte, color, s string) []byte {
	return append(b, h.colorize(color, s)...)
}

// stackOf — совместимость с github.com/pkg/errors и x/errors: стек отдельной строкой,
// как %wEx в Boot.
func stackOf(v any) string {
	if st, ok := v.(interface{ StackTrace() []byte }); ok {
		return string(st.StackTrace())
	}
	return ""
}

// --------------------------------------------------------------- подсобное

func lvlName(l slog.Level) string {
	switch {
	case l >= slog.LevelError:
		return "ERROR"
	case l >= slog.LevelWarn:
		return "WARN"
	case l >= slog.LevelInfo:
		return "INFO"
	case l >= slog.LevelDebug:
		return "DEBUG"
	default:
		return "TRACE"
	}
}

func lvlColor(l slog.Level) string {
	switch {
	case l >= slog.LevelError:
		return red
	case l >= slog.LevelWarn:
		return yellow
	case l >= slog.LevelInfo:
		return green
	default:
		return faint
	}
}

func msgColor(l slog.Level) string {
	if l >= slog.LevelError {
		return red
	}
	return ""
}

func asString(v slog.Value) string { return fmt.Sprintf("%v", v.Resolve()) }

// %15.15t — обрезка по первым N символам, затем дополняется слева.
func trunc(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func padLeft(s string, n int) string  { return strings.Repeat(" ", max(0, n-len(s))) + s }
func padRight(s string, n int) string { return s + strings.Repeat(" ", max(0, n-len(s))) }

// abbrev — logback logger{N}: все сегменты, кроме последнего, сворачиваются в одну букву.
func abbrev(s string, n int) string {
	if len(s) <= n {
		return s
	}
	sep := "."
	if strings.Contains(s, "/") {
		sep = "/"
	}
	parts := strings.Split(s, sep)
	if len(parts) < 2 {
		return trunc(s, n)
	}
	for i := range len(parts) - 1 {
		if parts[i] != "" {
			parts[i] = parts[i][:1]
		}
	}
	return trunc(strings.Join(parts, sep), n)
}

func bytesTrimSpace(b []byte) []byte {
	for len(b) > 0 && (b[len(b)-1] == '\n' || b[len(b)-1] == ' ') {
		b = b[:len(b)-1]
	}
	return b
}

// ------------------------------------------------------------- ротация файла

// rotate — size-based ротация без внешних зависимостей.
// Файл закрывается ДО переименования: на Windows переименовать файл,
// открытый текущим процессом, нельзя без FILE_SHARE_DELETE.
type rotate struct {
	path string
	max  int64
	keep int

	mu sync.Mutex
	f  *os.File
	n  int64
}

func newRotate(path string, mb, keep int) (*rotate, error) {
	if d := filepath.Dir(path); d != "" && d != "." {
		if err := os.MkdirAll(d, 0o755); err != nil {
			return nil, fmt.Errorf("log: mkdir %s: %w", d, err)
		}
	}
	r := &rotate{path: path, max: int64(mb) << 20, keep: max(1, keep)}
	return r, r.open()
}

func (r *rotate) open() error {
	f, err := os.OpenFile(r.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return fmt.Errorf("log: open %s: %w", r.path, err)
	}
	fi, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return fmt.Errorf("log: stat %s: %w", r.path, err)
	}
	r.f, r.n = f, fi.Size()
	return nil
}

func (r *rotate) Write(p []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.max > 0 && r.n+int64(len(p)) > r.max {
		if err := r.shift(); err != nil {
			return 0, err
		}
	}
	n, err := r.f.Write(p)
	r.n += int64(n)
	return n, err
}

// shift: app.jsonl.{keep-1} -> .{keep}, ..., .1 -> .2, app.jsonl -> .1
func (r *rotate) shift() error {
	if err := r.f.Close(); err != nil {
		return err
	}
	for i := r.keep - 1; i >= 1; i-- {
		from := fmt.Sprintf("%s.%d", r.path, i)
		to := fmt.Sprintf("%s.%d", r.path, i+1)
		if err := os.Rename(from, to); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("log: rotate %s -> %s: %w", from, to, err)
		}
	}
	if err := os.Rename(r.path, r.path+".1"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("log: rotate %s: %w", r.path, err)
	}
	return r.open()
}

func (r *rotate) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.f == nil {
		return nil
	}
	return r.f.Close()
}
