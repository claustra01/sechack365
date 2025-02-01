package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

func bindArticle(a *model.ArticleWithUser) openapi.Article {
	article := openapi.Article{
		Id:      a.Id,
		Title:   a.Title,
		Content: a.Content,
		User: openapi.SimpleUser{
			Username:    a.User.Username,
			Protocol:    a.User.Protocol,
			DisplayName: a.User.DisplayName,
			Icon:        a.User.Icon,
		},
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
	return article
}

func bindArticleComment(c *model.ArticleCommentWithUser) openapi.ArticleComment {
	comment := openapi.ArticleComment{
		Id:      c.Id,
		Content: c.Content,
		User: openapi.SimpleUser{
			Username:    c.User.Username,
			Protocol:    c.User.Protocol,
			DisplayName: c.User.DisplayName,
			Icon:        c.User.Icon,
		},
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
	return comment
}

func CreateArticle(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var articleRequestBody openapi.NewArticle
		body, err := io.ReadAll(r.Body)
		if err != nil {
			c.Logger.Warn("Failed to read request body", "Error", cerror.Wrap(err, "failed to create article"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(body, &articleRequestBody); err != nil {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to parse JSON"))
			returnError(w, http.StatusBadRequest)
			return
		}

		// validate content
		if articleRequestBody.Title == "" || articleRequestBody.Content == "" || len(articleRequestBody.Title) > 100 {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrEmptyContent, "failed to create article"))
			returnError(w, http.StatusBadRequest)
			return
		}

		// get current user
		user, err := c.CurrentUser(r)
		if err != nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to create article"))
			returnError(w, http.StatusUnauthorized)
			return
		}

		// create article
		articleId := util.NewUuid().String()
		if err := c.Controllers.Article.Create(articleId, user.Id, articleRequestBody.Title, articleRequestBody.Content); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create article"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		// create post
		postId := util.NewUuid().String()
		postContent := fmt.Sprintf("記事を投稿しました！\nhttps://%s/article/%s", c.Config.Host, articleId)

		// post to activitypub
		keyId := c.Controllers.ActivityPub.NewKeyIdUrl(user.Identifiers.Activitypub.Host, user.Id)
		privKeyPem, err := c.Controllers.User.GetActivityPubPrivKey(user.Id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		apPrivKey, _, err := util.DecodePem(privKeyPem)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		followers, err := c.Controllers.Follow.FindActivityPubRemoteFollowers(user.Id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		for _, follower := range followers {
			f := strings.Split(follower, "@")
			targetUrl, err := c.Controllers.ActivityPub.ResolveWebfinger(f[0], f[1])
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			targetActor, err := c.Controllers.ActivityPub.ResolveRemoteActor(targetUrl)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			activity := &model.ApActivity{
				Context: *c.Controllers.ActivityPub.NewApContext(),
				Type:    model.ActivityTypeCreate,
				Id:      fmt.Sprintf("https://%s/posts/%s", user.Identifiers.Activitypub.Host, postId),
				Actor:   c.Controllers.ActivityPub.NewActorUrl(user.Identifiers.Activitypub.Host, user.Id),
				Object: &model.ApNoteActivity{
					Id:        fmt.Sprintf("https://%s/posts/%s", user.Identifiers.Activitypub.Host, postId),
					Type:      model.ActivityTypeNote,
					Content:   util.WrapURLWithAnchor(postContent),
					Published: time.Now(),
				},
			}
			if _, err := c.Controllers.ActivityPub.SendActivity(keyId, apPrivKey, targetActor.Inbox, activity); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
				returnError(w, http.StatusInternalServerError)
				return
			}
		}

		// post to nostr
		nostrPrivKey, err := c.Controllers.User.GetNostrPrivKey(user.Id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if err := c.Controllers.Nostr.PublishPost(nostrPrivKey, postContent); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		// create
		if err := c.Controllers.Post.Create(postId, user.Id, postContent); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if err := c.Controllers.Article.CreateArticlePostRelation(articleId, postId); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create article post relation"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		returnResponse(w, http.StatusCreated, ContentTypeJson, nil)
	}
}

func GetArticle(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrEmptyContent, "failed to get article"))
			returnError(w, http.StatusBadRequest)
			return
		}

		article, err := c.Controllers.Article.FindById(id)
		if err != nil {
			c.Logger.Warn("Not Found", "Error", cerror.Wrap(err, "failed to get article"))
			returnError(w, http.StatusNotFound)
			return
		}

		returnResponse(w, http.StatusOK, ContentTypeJson, bindArticle(article))
	}
}

func GetArticleCommentsById(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrEmptyContent, "failed to get article comments"))
			returnError(w, http.StatusBadRequest)
			return
		}

		comments, err := c.Controllers.Article.FindCommentsByArticleId(id)
		if err != nil {
			c.Logger.Warn("Not Found", "Error", cerror.Wrap(err, "failed to get article comments"))
			returnError(w, http.StatusNotFound)
			return
		}

		var commentList []openapi.ArticleComment
		for _, comment := range comments {
			commentList = append(commentList, bindArticleComment(comment))
		}

		returnResponse(w, http.StatusOK, ContentTypeJson, commentList)
	}
}
