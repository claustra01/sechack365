package handler

import (
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func GenerateMock(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]interface{}{
			"message": "Hello, Developer!",
		})
	}
}
