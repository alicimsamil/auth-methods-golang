package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		authHeader := request.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(rw, "missing authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.Split(authHeader, " ")

		if len(token) != 2 {
			http.Error(rw, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		if err := verifyToken(token[1]); err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
		}

		handlerFunc.ServeHTTP(rw, request)
	}
}

func verifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil || !parsedToken.Valid {
		return errors.New("invalid token")
	}

	return nil
}
