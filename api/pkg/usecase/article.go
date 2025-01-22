package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IArticleRepository interface {
	Create(id, userId, title, content string) error
	FindById(id string) (*model.ArticleWithUser, error)
	CreateArticleComment(articleId, userId, content string) error
	FindCommentsByArticleId(articleId string) ([]*model.ArticleCommentWithUser, error)
}

type ArticleUsecase struct {
	ArticleRepository IArticleRepository
}

func (u *ArticleUsecase) Create(id, userId, title, content string) error {
	return u.ArticleRepository.Create(id, userId, title, content)
}

func (u *ArticleUsecase) FindById(id string) (*model.ArticleWithUser, error) {
	return u.ArticleRepository.FindById(id)
}

func (u *ArticleUsecase) CreateArticleComment(articleId, userId, content string) error {
	return u.ArticleRepository.CreateArticleComment(articleId, userId, content)
}

func (u *ArticleUsecase) FindCommentsByArticleId(articleId string) ([]*model.ArticleCommentWithUser, error) {
	return u.ArticleRepository.FindCommentsByArticleId(articleId)
}
