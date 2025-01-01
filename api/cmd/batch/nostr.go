package batch

import (
	"time"

	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
)

func UpdateNostrRemotePosts(c *framework.Context) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			// get all following nostr public keys
			pubKeys, err := c.Controllers.User.GetAllFollowingNostrPubKeys()
			if err != nil {
				c.Logger.Error("failed to update remote nostr posts", err)
				continue
			}

			// get latest remote post in local db
			latest, err := c.Controllers.Post.GetLatestNostrRemotePost()
			if err != nil {
				c.Logger.Error("failed to update remote nostr posts", err)
				continue
			}
			if latest == nil {
				latest = &model.Post{
					CreatedAt: time.Now().Add(-5 * time.Minute),
				}
			}

			// get new remote posts
			events, err := c.Controllers.Nostr.GetRemotePosts(pubKeys, latest.CreatedAt)
			if err != nil {
				c.Logger.Error("failed to update remote nostr posts", err)
				continue
			}
			if len(events) == 0 {
				continue
			}

			// save remote posts in local db
			if err := c.Controllers.Post.InsertNostrRemotePosts(events); err != nil {
				c.Logger.Error("failed to update remote nostr posts", err)
				continue
			}
			c.Logger.Info("remote nostr posts updated")
		}
	}()
}
