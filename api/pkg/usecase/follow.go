package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IFollowRepository interface {
	Create(followerId, followeeId string) (*model.Follow, error)
	UpdateAcceptFollow(followerId, followeeId string) (*model.Follow, error)
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
