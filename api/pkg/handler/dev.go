package handler

import (
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
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
		defaultIcon := fmt.Sprintf("https://%s/static/default_icon.png", c.Config.Host)
		if err := c.Controllers.User.CreateLocalUser("mock", "password", "Mock User", "This is mock user", defaultIcon, c.Config.Host); err != nil {
			panic(err)
		}
		if err := c.Controllers.NostrRelay.Create("wss://yabu.me"); err != nil {
			panic(err)
		}

		user, _ := c.Controllers.User.FindByLocalUsername("mock")
		privKey, _ := c.Controllers.User.GetNostrPrivKey(user.Id)
		profile := &model.NostrProfile{
			Name:        "mock",
			DisplayName: "Mock User",
			About:       "This is mock user",
			Picture:     defaultIcon,
		}
		if err := c.Controllers.Nostr.PublishProfile(privKey, profile); err != nil {
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
		posts, err := c.Controllers.Post.FindTimeline(0, 10000)
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

		for _, post := range posts {
			if err := c.Controllers.Post.DeleteById(post.Id); err != nil {
				panic(err)
			}
		}
		for _, user := range users {
			if err := c.Controllers.User.DeleteById(user.Id); err != nil {
				panic(err)
			}
		}

		nostrRelays, err := c.Controllers.NostrRelay.FindAll()
		if err != nil {
			panic(err)
		}
		for _, nostrRelay := range nostrRelays {
			if err := c.Controllers.NostrRelay.Delete(nostrRelay.Id); err != nil {
				panic(err)
			}
		}

		if _, err := w.Write([]byte("Mock Data Deleted")); err != nil {
			// NOTE: err should be nil
			panic(err)
		}
	}
}
