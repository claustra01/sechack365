package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IFollowRepository interface {
	Insert(follower string, followee string) (*model.Follow, error)
}

type FollowUsecase struct {
	FollowRepository IFollowRepository
}

func (u *FollowUsecase) Insert(follower string, followee string) (*model.Follow, error) {
	return u.FollowRepository.Insert(follower, followee)
}
