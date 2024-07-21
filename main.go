package main

import (
	"auth-methods/basic"
	"auth-methods/session"
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

	//Basic authentication
	router.HandleFunc("/basic-main", basic.BasicAuthMiddleWare(greetings))

	//Token based authentication
	router.HandleFunc("/token-main", token.JWTAuthMiddleware(greetings))
	router.HandleFunc("/jwt-login", token.UserLoginWithJWT).Methods(http.MethodPost)

	//Session based authentication
	router.HandleFunc("/session-main", session.SessionAuthMiddleware(greetings))
	router.HandleFunc("/session-login", session.UserLoginWithSession).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":80", router))
}
