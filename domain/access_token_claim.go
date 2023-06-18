package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type AccessTokenClaim struct {
	jwt.StandardClaims
	AccountID   int64  `json:"account_id"`
	AccountName string `json:"account_name"`
	Role        int    `json:"role"`
}
