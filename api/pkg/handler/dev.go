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
			if _, err := w.Write([]byte("Mock Data Already Exists")); err != nil {
				// NOTE: err should be nil
				panic(err)
			}
			return
		}
		if err := c.Controllers.Transaction.Begin(); err != nil {
			panic(err)
		}
		defer func() {
			if r := recover(); r != nil {
				if err := c.Controllers.Transaction.Rollback(); err != nil {
					panic(err)
				}
				panic(err)
			}
		}()
		user, err := c.Controllers.User.Create("mock", c.Config.Host, "local", "password", "Mock User", "This is mock user", "https://placehold.jp/150x150.png")
		if err != nil {
			panic(err)
		}
		if _, err := c.Controllers.ApUserIdentifier.Create(user.Id); err != nil {
			panic(err)
		}
		if err := c.Controllers.Transaction.Commit(); err != nil {
			panic(err)
		}
		if _, err := w.Write([]byte("Mock Data Created")); err != nil {
			// NOTE: err should be nil
			panic(err)
		}
	}
}

func ResetMock(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.Controllers.User.FindAll()
		if err != nil {
			panic(err)
		}
		if len(users) == 0 {
			if _, err := w.Write([]byte("Mock Data Not Found")); err != nil {
				// NOTE: err should be nil
				panic(err)
			}
			return
		}
		for _, user := range users {
			if err := c.Controllers.ApUserIdentifier.DeleteById(user.Id); err != nil {
				panic(err)
			}
			if err := c.Controllers.User.DeleteById(user.Id); err != nil {
				panic(err)
			}
		}
		if _, err := w.Write([]byte("Mock Data Deleted")); err != nil {
			// NOTE: err should be nil
			panic(err)
		}
	}
}
