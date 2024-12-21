package domain

import "github.com/golang-jwt/jwt"

type AuthToken struct {
	token *jwt.Token
}
