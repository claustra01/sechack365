package usecase

import (
	"github.com/claustra01/sechack365/pkg/model"
)

type IPostRepository interface {
	Create(userId, content string) error
	FindById(id string) (*model.PostWithUser, error)
	FindTimeline(offset int, limit int) ([]*model.PostWithUser, error)
	FindUserTimeline(userId string, offset int, limit int) ([]*model.PostWithUser, error)
	Delete(id string) error
}

type PostUsecase struct {
	PostRepository IPostRepository
}

func (u *PostUsecase) Create(userId, content string) error {
	return u.PostRepository.Create(userId, content)
}

func (u *PostUsecase) FindById(id string) (*model.PostWithUser, error) {
	return u.PostRepository.FindById(id)
}

func (u *PostUsecase) FindTimeline(offset int, limit int) ([]*model.PostWithUser, error) {
	return u.PostRepository.FindTimeline(offset, limit)
}

func (u *PostUsecase) FindUserTimeline(userId string, offset int, limit int) ([]*model.PostWithUser, error) {
	return u.PostRepository.FindUserTimeline(userId, offset, limit)
}

func (u *PostUsecase) Delete(id string) error {
	return u.PostRepository.Delete(id)
}
