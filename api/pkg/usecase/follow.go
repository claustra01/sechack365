package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IFollowRepository interface {
	Create(followerId, followeeId string) (*model.Follow, error)
	UpdateAcceptFollow(followerId, followeeId string) (*model.Follow, error)
	FindFollowsByUserId(userId string) ([]*model.User, error)
	FindFollowersByUserId(userId string) ([]*model.User, error)
}

type FollowUsecase struct {
	FollowRepository IFollowRepository
}

func (u *FollowUsecase) Create(followerId, followeeId string) (*model.Follow, error) {
	return u.FollowRepository.Create(followerId, followeeId)
}

func (u *FollowUsecase) UpdateAcceptFollow(followerId, followeeId string) (*model.Follow, error) {
	return u.FollowRepository.UpdateAcceptFollow(followerId, followeeId)
}

func (u *FollowUsecase) FindFollowsByUserId(userId string) ([]*model.User, error) {
	return u.FollowRepository.FindFollowsByUserId(userId)
}

func (u *FollowUsecase) FindFollowersByUserId(userId string) ([]*model.User, error) {
	return u.FollowRepository.FindFollowersByUserId(userId)
}
