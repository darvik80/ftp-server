package account

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"ftp-server-go/internal/store"
	"ftp-server-go/internal/store/repository"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrAuthFailed        = errors.New("authentication failed")
)

type Account struct {
	Username string
	Home     string
	CanWrite bool
	CanRead  bool
	CanList  bool
}

type Accounts interface {
	Register(ctx context.Context, username, password string) error
	Exists(ctx context.Context, username string) (bool, error)
	Authenticate(ctx context.Context, username, password string) (*Account, error)
}

type acc struct {
	account *repository.AccountRepository
}

func New(account *repository.AccountRepository) Accounts {
	return &acc{
		account: account,
	}
}

func (a *acc) Exists(ctx context.Context, username string) (bool, error) {
	return a.account.Exists(ctx, username)
}

func (a *acc) Authenticate(ctx context.Context, username, password string) (*Account, error) {
	account, err := a.account.FindAccount(ctx, username)
	if err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if !verifyPassword(password, account.Password) {
		return nil, ErrAuthFailed
	}

	return &Account{
		Username: account.Username,
		CanWrite: true,
		CanRead:  true,
		CanList:  true,
	}, nil
}

func (a *acc) Register(ctx context.Context, username, password string) error {
	hash, err := hashPassword(password)
	if err != nil {
		return err
	}
	if err := a.account.CreateAccount(ctx, username, hash); err != nil {
		if errors.Is(err, store.ErrRecordAlreadyExists) {
			return ErrUserAlreadyExists
		}
		return err
	}

	return nil
}

const (
	time    = 3
	memory  = 64 * 1024 // KiB
	threads = 4
	keyLen  = 32
	saltLen = 16
)

func hashPassword(pw string) (string, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	dk := argon2.IDKey([]byte(pw), salt, time, memory, threads, keyLen)
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, time, threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(dk)), nil
}

func verifyPassword(pw, encoded string) bool {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 {
		return false
	}
	var m uint32
	var t uint8
	var p uint8
	fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &m, &t, &p) // парси аккуратно, это скелет
	salt, _ := base64.RawStdEncoding.DecodeString(parts[4])
	want, _ := base64.RawStdEncoding.DecodeString(parts[5])
	got := argon2.IDKey([]byte(pw), salt, uint32(t), m, p, uint32(len(want)))
	return subtle.ConstantTimeCompare(got, want) == 1
}
