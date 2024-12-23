package service

import (
	"github.com/dmosyan/Learning-Go/apis/banking-auth/dto"
	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	Verify(urlParams map[string]string) *errs.AppError
	Refresh(request dto.RefreshTokenRequest) (*dto.LoginResponse, *errs.AppError)
}
