package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

var (
	// ErrLineTooLong — содержимое длиннее протокольного лимита, LF не пришёл.
	ErrLineTooLong = errors.New("line too long")
	// ErrTruncated — источник умер на строке без терминатора (обрыв, не QUIT).
	ErrTruncated = errors.New("line truncated")
	ErrIdle      = errors.New("idle too long")
)

const (
	emptyString = ""
	maxIdle     = 10
)

type processor struct {
	r   io.Reader
	buf []byte
	off int

	readErr error
	idle    int
}

func NewProcessor(size int, r io.Reader) *processor {
	if size <= 2 {
		size = 1024
	}
	return &processor{
		buf: make([]byte, size),
		r:   r,
	}
}

func (p *processor) ReadString() (string, error) {
	// extract string from buffer
	get := func() (string, bool) {
		if idx := bytes.IndexByte(p.buf[:p.off], '\n'); idx >= 0 {
			end := idx
			if end > 0 && p.buf[end-1] == '\r' {
				end--
			}
			str := string(p.buf[:end])

			rest := p.off - idx - 1
			copy(p.buf, p.buf[idx+1:])
			p.off = rest

			return str, true
		}

		return emptyString, false
	}

	// fail and reset processor
	fail := func(err error) error {
		p.off, p.readErr, p.idle = 0, nil, 0
		return err
	}

	for {
		if res, ok := get(); ok {
			return res, nil
		}

		if p.off >= len(p.buf) {
			return emptyString, fail(ErrLineTooLong)
		}

		if p.readErr != nil {
			err := p.readErr
			if p.off > 0 {
				return emptyString, fail(ErrTruncated)
			}

			return emptyString, fail(err)
		}

		n, err := p.r.Read(p.buf[p.off:])
		p.off += n

		// network close but we still could have some data
		if err != nil {
			p.readErr = err
			continue
		}

		// too much empty data
		if n == 0 {
			if p.idle++; p.idle == maxIdle {
				return emptyString, fail(ErrIdle)
			}
			continue
		}

		p.idle = 0

	}
}

func main() {
	var rw io.ReadWriter = new(bytes.Buffer)
	p := NewProcessor(16, rw)

	rw.Write([]byte("hello world\r\r\n"))
	for {
		if s, err := p.ReadString(); err == nil {
			fmt.Println(s)
		} else {
			break
		}
	}
	fmt.Println("===")
	rw.Write([]byte("\n\ntest\nRSET\nSTOR\n"))
	for {
		if s, err := p.ReadString(); err == nil {
			fmt.Println(s)
		} else {
			break
		}
	}
}
