package verify

import (
	"net/http"
	"purple_school_golang_advanced/3_validation_api/configs"
)

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux, conf *configs.AppConfig) {
	handler := &VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send(*conf))
	router.HandleFunc("GET /verify/{code}", handler.Verify())
}

func (handler *VerifyHandler) Send(conf configs.AppConfig) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		sendQuery := req.URL.Query()
		login := sendQuery.Get("login")
		if login != "" {
			writer.WriteHeader(200)
			writer.Write([]byte("SEND"))

			SendVerify(login, conf.MailConfig)

		} else {
			writer.WriteHeader(402)
			writer.Write([]byte("Не верные данные"))
		}
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		code := req.PathValue("code")
		if code != "" {
			if VerifyHash(code) {
				writer.WriteHeader(200)
				writer.Write([]byte("Код проверен и соответсвует пользователю"))
			} else {
				writer.WriteHeader(403)
				writer.Write([]byte("Error"))
			}
		} else {
			writer.WriteHeader(403)
		}
	}
}
