package main

import (
	"context"
	"errors"
	"fmt"
	"ftp-server-go/internal/config"
	"ftp-server-go/internal/log"
	"ftp-server-go/internal/service/engine"
	"ftp-server-go/internal/service/engine/account"
	"ftp-server-go/internal/service/server"
	"ftp-server-go/internal/store"
	"ftp-server-go/internal/store/repository"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func run(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logger, closeLog, err := log.Setup(cfg.Logger)
	if err != nil {
		return fmt.Errorf("setup logger: %w", err)
	}
	defer func() { _ = closeLog() }()

	ctx = log.WithLogger(ctx, logger)

	logger.Info("starting server", "version", version, "commit", commit, "date", date)
	g, gCtx := errgroup.WithContext(ctx)

	db, err := store.Open(ctx, cfg)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	accRepo := repository.NewAccountRepository(db)
	if err := repository.AutoMigrate(ctx, db); err != nil {
		return fmt.Errorf("auto migrate: %w", err)
	}
	acc := account.New(accRepo)
	for _, a := range cfg.Database.Accounts {
		if err := acc.Register(ctx, a.Username, a.Password); err != nil {
			return fmt.Errorf("create account: %w", err)
		}
	}

	e := engine.NewDefaultEngine(acc)
	e.RegisterCommands()
	s := server.NewServer(cfg, e)

	g.Go(func() error {
		return s.Run(gCtx)
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		logger.Error("services stopped with errors", err)
		os.Exit(1)
	}

	return nil
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
