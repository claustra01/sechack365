package main

import (
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func SetupRouter(r *framework.Router) error {
	// TODO: check error
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	r.POST("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	r.HandleRoutes()
	return nil
}
