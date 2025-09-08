package main

import (
	"go_example/internal/auth"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()

}
