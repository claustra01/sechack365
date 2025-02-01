package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IArticleRepository interface {
	Create(id, userId, title, content string) error
	FindById(id string) (*model.ArticleWithUser, error)
	CreateArticleComment(articleId, userId, content string) error
	FindCommentsByArticleId(articleId string) ([]*model.ArticleCommentWithUser, error)
	// Article連携
	CreateArticlePostRelation(articleId, postId string) error
	FindArticlePostRelation(postId string) (*model.ArticlePostRelation, error)
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

func (u *ArticleUsecase) CreateArticlePostRelation(articleId, postId string) error {
	return u.ArticleRepository.CreateArticlePostRelation(articleId, postId)
}

func (u *ArticleUsecase) FindArticlePostRelation(postId string) (*model.ArticlePostRelation, error) {
	return u.ArticleRepository.FindArticlePostRelation(postId)
}
