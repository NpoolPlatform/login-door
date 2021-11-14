package response

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"message": msg})
}

func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload) //nolint: errcheck
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response) //nolint: errcheck
}
