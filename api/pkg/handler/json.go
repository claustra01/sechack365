package handler

import (
	"encoding/json"
	"net/http"

	"github.com/claustra01/sechack365/pkg/openapi"
)

const (
	ContentTypeJson   = "application/json"
	ContentTypeApJson = "application/activity+json"
)

var errorResponse = map[int]any{
	http.StatusBadRequest: openapi.Error400{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad Request",
	},
	http.StatusUnauthorized: openapi.Error401{
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized",
	},
	http.StatusNotFound: openapi.Error404{
		StatusCode: http.StatusNotFound,
		Message:    "Not Found",
	},
	http.StatusConflict: openapi.Error409{
		StatusCode: http.StatusConflict,
		Message:    "Conflict",
	},
	http.StatusInternalServerError: openapi.Error500{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
	},
}

func returnResponse(w http.ResponseWriter, statusCode int, contentType string, responseBody any) {
	body, err := json.Marshal(responseBody)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	if _, err := w.Write(body); err != nil {
		// NOTE: err should be nil
		panic(err)
	}
}

func returnError(w http.ResponseWriter, statusCode int) {
	resp, ok := errorResponse[statusCode]
	if !ok {
		panic("unsupported status code")
	}
	body, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", ContentTypeJson)
	http.Error(w, string(body), statusCode)
}
