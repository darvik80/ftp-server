package repository

import (
	"context"
	"errors"
	"fmt"
	"ftp-server-go/internal/store"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique" json:"username"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type AccountRepository struct{ db *gorm.DB }

func NewAccountRepository(db *gorm.DB) *AccountRepository { return &AccountRepository{db: db} }

func (r *AccountRepository) FindAccount(ctx context.Context, username string) (*Account, error) {
	var account Account
	if err := r.db.WithContext(ctx).First(&account, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, store.ErrRecordNotFound
		}

		return nil, fmt.Errorf("find account: %w", err)
	}
	return &account, nil
}

func (r *AccountRepository) Exists(ctx context.Context, username string) (bool, error) {
	var account Account
	if err := r.db.WithContext(ctx).First(&account, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("check account existence: %w", err)
	}
	return true, nil
}

func (r *AccountRepository) CreateAccount(ctx context.Context, username, password string) error {
	if err := r.db.WithContext(ctx).Create(&Account{Username: username, Password: password}).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return store.ErrRecordAlreadyExists
		}
		return fmt.Errorf("create account: %w", err)
	}

	return nil
}

func AutoMigrate(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&Account{})
}
