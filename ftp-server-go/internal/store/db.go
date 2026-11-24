package store

import (
	"context"
	"errors"
	"fmt"
	"ftp-server-go/internal/config"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(ctx context.Context, cfg *config.Config) (*gorm.DB, error) {
	// Прагмы через DSN применяются до первого запроса, на каждое новое соединение.
	dsn := fmt.Sprintf(
		"file:%s?_journal_mode=WAL&_synchronous=NORMAL&_foreign_keys=ON&_busy_timeout=5000&_txlock=immediate",
		cfg.Database.Uri,
	)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true, // кэш подготовленных выражений: на частых однотипных запросах заметно
		NowFunc:     func() time.Time { return time.Now().UTC() },
		Logger: logger.New(
			log.New(log.Writer(), "\n", log.LstdFlags),
			logger.Config{SlowThreshold: 200 * time.Millisecond},
		),
		TranslateError: true, // gorm.ErrDuplicatedKey / ErrForeignKeyViolated вместо сырых ошибок драйвера
	})
	if err != nil {
		return nil, errors.New("open sqlite: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("open sqlite: " + err.Error())
	}
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		return nil, errors.New("ping: " + err.Error())
	}

	return db, nil
}
