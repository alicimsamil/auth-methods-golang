package main

import (
	token "auth-methods/token"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	greetings := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi there")
	}

	//router.HandleFunc("/", basic.BasicAuthMiddleWare(greetings))

	router.HandleFunc("/", token.JWTAuthMiddleware(greetings))
	router.HandleFunc("/jwt-login", token.UserLoginWithJWT).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":80", router))
}
