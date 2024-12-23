package domain

import (
	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/errs"
	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/logger"
	"github.com/golang-jwt/jwt"
)

const HMAC_SAMPLE_SECRET = "my_secret_key"

type AuthToken struct {
	token *jwt.Token
}

func (t AuthToken) NewAccessToken() (string, *errs.AppError) {
	signedString, err := t.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		logger.Error("failed while signing access token: " + err.Error())
		return "", errs.NewUnexpectedError("cannot generate access token")
	}
	return signedString, nil
}
