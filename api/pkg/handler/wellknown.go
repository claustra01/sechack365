package handler

import (
	"fmt"
	"net/http"
)

func Nodeinfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"version":"0.1.0"}`)
}

func Webfinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"subject":"acct:admin@localhost","links":[{"rel":"self","type":"application/activity+json","href":"http://localhost/api/v1/actor/admin"}]}`)
}
