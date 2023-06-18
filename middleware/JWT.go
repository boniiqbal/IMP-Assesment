package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type JSONFailed struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JWTClaims struct {
	jwt.StandardClaims
	AccountID   int64  `json:"account_id"`
	AccountName string `json:"account_name"`
	Role        int    `json:"role"`
}

// JWTVerify function to verify json web token
func JWTVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if os.Getenv("NO_TOKEN") == "1" {
			return next(c)
		}

		notAuth := []string{
			"/healthcheck",
			"/auth/login",
			"/auth/signup",
		} //List of endpoints that doesn't require auth
		requestPath := c.Request().URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath || strings.HasPrefix(requestPath, value) {
				return next(c)
			}
		}

		req := c.Request()
		header := req.Header
		auth := header.Get("Authorization")

		if len(auth) <= 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "authorization is empty")
		}

		splitToken := strings.Split(auth, " ")
		if len(splitToken) < 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "authorization is empty")
		}

		if splitToken[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "authorization is invalid")
		}

		tokenStr := splitToken[1]
		token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if claims, ok := token.Claims.(*JWTClaims); token.Valid && ok {
			c.Set("token", token)
			c.Set("accountID", claims.AccountID)
			c.Set("accountName", claims.AccountName)
			c.Set("role", claims.Role)

			return next(c)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			var errorStr string
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				errorStr = fmt.Sprintf("Invalid token format: %s", tokenStr)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				errorStr = "Token has been expired"
			} else {
				errorStr = fmt.Sprintf("Token Parsing Error: %s", err.Error())
			}
			return echo.NewHTTPError(http.StatusUnauthorized, errorStr)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unknown token error")
		}
	}
}
