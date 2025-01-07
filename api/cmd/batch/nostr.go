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
			pubKeys, err := c.Controllers.Follow.GetAllFollowingNostrPubKeys()
			if err != nil {
				c.Logger.Error("failed to update remote nostr posts", err)
				continue
			}
			if len(pubKeys) == 0 {
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

func UpdateNostrRemoteFollowers(c *framework.Context) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			// local user pubkeys
			pubKeys, err := c.Controllers.Follow.GetAllLocalUserNostrPubKeys()
			if err != nil {
				c.Logger.Error("failed to update remote nostr followers", err)
				continue
			}
			if len(pubKeys) == 0 {
				continue
			}

			// get latest follow
			latest, err := c.Controllers.Follow.GetLatestNostrRemoteFollow()
			if err != nil {
				c.Logger.Error("failed to update remote nostr followers", err)
				continue
			}
			if latest == nil {
				latest = &model.Follow{
					CreatedAt: time.Now().Add(-5 * time.Minute),
				}
			}

			// get new remote follows
			followers, err := c.Controllers.Nostr.GetRemoteFollowerPubKeys(pubKeys, latest.CreatedAt)
		}
	}()
}
