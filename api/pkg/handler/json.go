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
	if _, err := w.Write(body); err != nil {
		// NOTE: err should be nil
		panic(err)
	}
}

func jsonCustomContentTypeResponse(w http.ResponseWriter, data any, contentType string) {
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", contentType)
	if _, err := w.Write(body); err != nil {
		// NOTE: err should be nil
		panic(err)
	}
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

func returnBadRequest(w http.ResponseWriter, logger model.ILogger, errInput error) {
	logger.Warn("Bad Request", "Error", errInput)
	jsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
}

func returnNotFound(w http.ResponseWriter, logger model.ILogger, errInput error) {
	logger.Warn("Not Found", "Error", errInput)
	jsonErrorResponse(w, http.StatusNotFound, "Not Found")
}

func returnInternalServerError(w http.ResponseWriter, logger model.ILogger, errInput error) {
	logger.Error("Internal Server Error", "Error", errInput)
	jsonErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
}
