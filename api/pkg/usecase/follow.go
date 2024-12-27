package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IFollowRepository interface {
	Create(followerId, targetId string) error
	UpdateAcceptFollow(followerId, targetId string) error
	FindFollowsByUserId(userId string) ([]*model.User, error)
	FindFollowersByUserId(userId string) ([]*model.User, error)
}

type FollowUsecase struct {
	FollowRepository IFollowRepository
}

func (u *FollowUsecase) Create(followerId, targetId string) error {
	return u.FollowRepository.Create(followerId, targetId)
}

func (u *FollowUsecase) UpdateAcceptFollow(followerId, targetId string) error {
	return u.FollowRepository.UpdateAcceptFollow(followerId, targetId)
}

func (u *FollowUsecase) FindFollowsByUserId(userId string) ([]*model.User, error) {
	return u.FollowRepository.FindFollowsByUserId(userId)
}

func (u *FollowUsecase) FindFollowersByUserId(userId string) ([]*model.User, error) {
	return u.FollowRepository.FindFollowersByUserId(userId)
}
