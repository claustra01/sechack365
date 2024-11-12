package handler

import (
	"encoding/json"
	"net/http"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
)

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

func returnBadRequest(w http.ResponseWriter, logger model.ILogger, errInput error) {
	logger.Warn("Bad Request", "Error", errInput)
	resp := &openapi.Error400{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad Request",
	}
	body, err := json.Marshal(resp)
	if err != nil {
		// NOTE: err should be nil
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(body), http.StatusInternalServerError)
}

func returnUnauthorized(w http.ResponseWriter, logger model.ILogger, errInput error) {
	logger.Warn("Unauthorized", "Error", errInput)
	resp := &openapi.Error401{
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized",
	}
	body, err := json.Marshal(resp)
	if err != nil {
		// NOTE: err should be nil
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(body), http.StatusInternalServerError)
}

func returnNotFound(w http.ResponseWriter, logger model.ILogger, errInput error) {
	logger.Warn("Not Found", "Error", errInput)
	resp := &openapi.Error404{
		StatusCode: http.StatusNotFound,
		Message:    "Not Found",
	}
	body, err := json.Marshal(resp)
	if err != nil {
		// NOTE: err should be nil
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(body), http.StatusInternalServerError)
}

func returnInternalServerError(w http.ResponseWriter, logger model.ILogger, errInput error) {
	logger.Error("Internal Server Error", "Error", errInput)
	resp := &openapi.Error400{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
	}
	body, err := json.Marshal(resp)
	if err != nil {
		// NOTE: err should be nil
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(body), http.StatusInternalServerError)
}
