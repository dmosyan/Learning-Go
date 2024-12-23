package domain

import (
	"database/sql"

	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/errs"
	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/logger"
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

func (d AuthRepositoryDb) RefreshTokenExists(refreshToken string) *errs.AppError {
	sqlSelect := "SELECT refresh_token FROM refresh_token_store WHERE refresh_token = ?"
	var token string

	err := d.client.Get(&token, sqlSelect, refreshToken)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewAuthenticationError("refresh token is not registered in the store")
		} else {
			logger.Error("unexpected error from the database: " + err.Error())
			return errs.NewUnexpectedError("unexpected database error")
		}
	}
	return nil
}
