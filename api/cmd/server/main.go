package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "1323"
	}
	http.ListenAndServe(":"+port, nil)
}
