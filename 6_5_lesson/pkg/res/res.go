package res

import (
	"encoding/json"
	"net/http"
)

func Json(writer http.ResponseWriter, data any, statusCode int) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(data)
}
