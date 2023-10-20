package utils

import (
	"encoding/json"
	"net/http"
)

func SendHttpResponse(w http.ResponseWriter, status int, data interface{}) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		globalLogger.Error("error parsing json data", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(jsonResponse))
}
