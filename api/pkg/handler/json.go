package handler

import (
	"encoding/json"
	"net/http"

	"github.com/claustra01/sechack365/pkg/model"
)

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func jsonResponse(w http.ResponseWriter, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func jsonErrorResponse(w http.ResponseWriter, code int, message string) {
	resp := &ErrorResponse{
		StatusCode: code,
		Message:    message,
	}
	body, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(body), http.StatusInternalServerError)
}

func returnInternalServerError(w http.ResponseWriter, logger model.ILogger, errInput error) {
	logger.Error("Internal Server Error:", "Error", errInput)
	jsonErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
}
