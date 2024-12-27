package usecase

import (
	"time"

	"github.com/claustra01/sechack365/pkg/model"
)

type IPostRepository interface {
	Create(userId, content string) error
	FindById(id string) (*model.PostWithUser, error)
	FindTimeline(createdAt time.Time, limit int) ([]*model.PostWithUser, error)
	FindUserTimeline(userId string, createdAt time.Time, limit int) ([]*model.PostWithUser, error)
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

func (u *PostUsecase) FindTimeline(createdAt time.Time, limit int) ([]*model.PostWithUser, error) {
	return u.PostRepository.FindTimeline(createdAt, limit)
}

func (u *PostUsecase) FindUserTimeline(userId string, createdAt time.Time, limit int) ([]*model.PostWithUser, error) {
	return u.PostRepository.FindUserTimeline(userId, createdAt, limit)
}

func (u *PostUsecase) Delete(id string) error {
	return u.PostRepository.Delete(id)
}
