package verify

import (
	"net/http"
	"purple_school_golang_advanced/3_validation_api/configs"
	"purple_school_golang_advanced/3_validation_api/pkg/req"
)

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux, conf *configs.AppConfig) {
	handler := &VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send(*conf))
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send(conf configs.AppConfig) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequest](&writer, r)
		if err != nil {
			return
		}
		resp := SendResponce{
			Result: "ok",
		}
		SendVerify(body.Email, conf.MailConfig)
		req.Json(writer, resp, 200)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		resp := VerifyResponce{
			Result: "ok",
		}

		if hash == "" {
			resp.Result = "error"
			req.Json(writer, resp, 402)
			return
		}

		if !VerifyHash(hash) {
			resp.Result = "error"
			req.Json(writer, resp, 403)
			return
		}

		DeleteHash(hash)
		req.Json(writer, resp, 200)
	}
}
