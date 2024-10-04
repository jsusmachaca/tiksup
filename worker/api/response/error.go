package response

import (
	"encoding/json"
	"net/http"
)

func WriteJsonError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	respone := map[string]string{
		"error": message,
	}
	if err := json.NewEncoder(w).Encode(respone); err != nil {
		http.Error(w, message, statusCode)
	}
}
