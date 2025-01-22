package handler

import (
	"encoding/json"
	"io"
	"net/http"

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
		uuid := util.NewUuid().String()
		if err := c.Controllers.Article.Create(uuid, user.Id, articleRequestBody.Title, articleRequestBody.Content); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create article"))
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
