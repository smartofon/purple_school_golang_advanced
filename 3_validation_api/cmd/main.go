package main

import (
	"net/http"
	"purple_school_golang_advanced/3_validation_api/configs"
	"purple_school_golang_advanced/3_validation_api/internal/storage"
	"purple_school_golang_advanced/3_validation_api/pkg/verify"
)

var AppConf = &configs.AppConfig{}

func main() {
	// загружаем параметры конфигурации приложения
	AppConf.LoadConfig()

	// Инициализация глобального хранилища (вызывается автоматически при старте)
	storage.GlobalStorage = &storage.VerifyStorage{
		Data: make(map[string]string),
		File: "storage/verify.json",
	}

	_ = storage.GlobalStorage.Load()

	router := http.NewServeMux()
	verify.NewVerifyHandler(router, AppConf)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
