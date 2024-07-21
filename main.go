package main

import (
	"auth-methods/middleware/basic"
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

	router.HandleFunc("/", basic.BasicAuthMiddleWare(greetings))

	log.Fatal(http.ListenAndServe(":80", router))
}
