package token

import (
	"auth-methods/common"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func UserLoginWithJWT(rw http.ResponseWriter, r *http.Request) {
	user, err := common.DecodeUser(r)
	if err != nil {
		http.Error(rw, "entity is not processable", http.StatusUnprocessableEntity)
		return
	}

	// Perform login actions with user here.

	token, err := createJWTToken(user.Email)
	if err != nil {
		http.Error(rw, "could not create token", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func createJWTToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 45).Unix(),
	})

	tokenStr, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
