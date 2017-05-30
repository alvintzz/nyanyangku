package middleware

import (
	"net/http"
	"time"
	"encoding/json"
)

type Response struct {
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"message_error"`
	ProcessTime  string      `json:"process_time"`
	Data         interface{} `json:"data,omitempty"`
}

type HandlerJSON func(rw http.ResponseWriter, r *http.Request) (interface{}, error)
func (fn HandlerJSON) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := Response{}

	w.Header().Set("Content-Type", "application/json")
	start := time.Now()

	data, err := fn(w, r)
	if err != nil {
		response.ErrorMessage = err.Error()
	} else {
		response.Success = true
		response.ErrorMessage = "Success"
		response.Data = data
	}
	response.ProcessTime = time.Since(start).String()

	responseJSON, err := json.Marshal(response)
	if err == nil {
		w.Write(responseJSON)
		return
	}

	response = Response{
		Success: false,
		ErrorMessage: "Terjadi kesalahan pada server. Cobalah beberapa saat lagi.",
	}
	responseJSON, _ = json.Marshal(response)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(responseJSON)
}
