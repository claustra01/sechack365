package controller

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/repository"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type ArticleController struct {
	ArticleUsecase usecase.ArticleUsecase
}

func NewArticleController(conn model.ISqlHandler) *ArticleController {
	return &ArticleController{
		ArticleUsecase: usecase.ArticleUsecase{
			ArticleRepository: &repository.ArticleRepository{
				SqlHandler: conn,
			},
		},
	}
}

func (c *ArticleController) Create(id, userId, title, content string) error {
	return c.ArticleUsecase.Create(id, userId, title, content)
}

func (c *ArticleController) FindById(id string) (*model.ArticleWithUser, error) {
	return c.ArticleUsecase.FindById(id)
}

func (c *ArticleController) CreateArticleComment(articleId, userId, content string) error {
	return c.ArticleUsecase.CreateArticleComment(articleId, userId, content)
}

func (c *ArticleController) FindCommentsByArticleId(articleId string) ([]*model.ArticleCommentWithUser, error) {
	return c.ArticleUsecase.FindCommentsByArticleId(articleId)
}

func (c *ArticleController) CreateArticlePostRelation(articleId, postId string) error {
	return c.ArticleUsecase.CreateArticlePostRelation(articleId, postId)
}

func (c *ArticleController) FindArticlePostRelation(postId string) (*model.ArticlePostRelation, error) {
	return c.ArticleUsecase.FindArticlePostRelation(postId)
}
