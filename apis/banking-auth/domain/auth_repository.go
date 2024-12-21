package domain

import (
	"github.com/dmosyan/Learning-Go/apis/banking/errs"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *errs.AppError)
	GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *errs.AppError)
	RefreshTokenExists(refreshToken string) *errs.AppError
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func NewAuthRepository(c *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client: c}
}
