package main

import "net/http"

type MyHandler struct{}

func NewMyHandler(router *http.ServeMux) {
	handler := &MyHandler{}
	router.HandleFunc("/ping", handler.ping())
}

func (handler *MyHandler) ping() http.HandlerFunc {
	return func(whriter http.ResponseWriter, req *http.Request) {
		whriter.Write([]byte("ping ok version 2"))
	}
}
