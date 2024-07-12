package main

import (
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func SetupRouter() *framework.Router {
	r := framework.NewRouter()

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	r.POST("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	return r
}
