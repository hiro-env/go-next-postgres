package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// 本来Secrets Manager KeyVaultなどからSDKで取得する
var jwtSecretKey = []byte("jwt_secret_key")

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		if path == "/login" || path == "/register" || path == "/logout" {
			return next(c)
		}

		cookie, err := c.Cookie("token")
		if err != nil {
			return c.String(http.StatusUnauthorized, "authorization failed")
		}

		tokenString := cookie.Value

		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})
		if err != nil || !token.Valid {
			return c.String(http.StatusUnauthorized, "authorization failed")
		}

		claims := token.Claims.(*jwt.StandardClaims)
		userID, err := strconv.Atoi(claims.Id)
		if err != nil {
			return c.String(http.StatusUnauthorized, "authorization failed")
		}

		c.Set("userID", int64(userID))

		return next(c)
	}
}
