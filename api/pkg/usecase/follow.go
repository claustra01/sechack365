package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IFollowRepository interface {
	Insert(followerId string, followeeId string) (*model.Follow, error)
	UpdateAcceptFollow(followerId string, followeeId string) (*model.Follow, error)
}

type FollowUsecase struct {
	FollowRepository IFollowRepository
}

func (u *FollowUsecase) Insert(followerId string, followeeId string) (*model.Follow, error) {
	return u.FollowRepository.Insert(followerId, followeeId)
}

func (u *FollowUsecase) UpdateAcceptFollow(followerId string, followeeId string) (*model.Follow, error) {
	return u.FollowRepository.UpdateAcceptFollow(followerId, followeeId)
}
