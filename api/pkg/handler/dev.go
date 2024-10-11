package handler

import (
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func GenerateMock(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.Controllers.User.FindAll()
		if err != nil {
			panic(err)
		}
		if len(users) > 0 {
			w.Write([]byte("Mock Data Already Exists"))
			return
		}
		if _, err := c.Controllers.User.Create("mock", c.Config.Host, "local", "Mock User", "This is mock user", "https://placehold.jp/150x150.png"); err != nil {
			panic(err)
		}
		w.Write([]byte("Mock Data Created"))
	}
}
