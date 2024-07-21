package token

import (
	"auth-methods/request"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func UserLoginWithJWT(rw http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
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

func decodeUser(req *http.Request) (request.UserRequest, error) {
	var user request.UserRequest
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil || user.Email == "" || user.Password == "" {
		return user, errors.New("invalid")
	}
	return user, nil
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
