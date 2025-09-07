package main

import (
	"fmt"
	"net/http"
)

func ping(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("ping ok"))
	fmt.Println(request)
}

func main() {

	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8081", nil)
}
