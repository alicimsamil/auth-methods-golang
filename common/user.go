package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

func DecodeUser(req *http.Request) (UserRequest, error) {
	var user UserRequest
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil || user.Email == "" || user.Password == "" {
		return user, errors.New("invalid")
	}
	return user, nil
}
