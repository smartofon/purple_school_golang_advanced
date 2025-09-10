package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

func ping(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(strconv.Itoa(rand.IntN(6) + 1)))
}

func main() {
	http.HandleFunc("/", ping)
	http.ListenAndServe(":8081", nil)
}
