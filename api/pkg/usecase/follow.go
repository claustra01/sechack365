package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IFollowRepository interface {
	Create(followerId, targetId string) error
	UpdateAcceptFollow(followerId, targetId string) error
	FindFollowByFollowerAndTarget(followerId, targetId string) (*model.Follow, error)
	FindFollowsByUserId(userId string) ([]*model.SimpleUser, error)
	FindFollowersByUserId(userId string) ([]*model.SimpleUser, error)
	FindNostrFollowPublicKeys(userId string) ([]string, error)
	CheckIsFollowing(followerId, targetId string) (bool, error)
	Delete(followerId, targetId string) error
	// Nostr用
	GetAllFollowingNostrPubKeys() ([]string, error)
	GetAllLocalUserNostrPubKeys() ([]string, error)
	GetLatestNostrRemoteFollow() (*model.Follow, error)
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

func (u *FollowUsecase) FindFollowByFollowerAndTarget(followerId, targetId string) (*model.Follow, error) {
	return u.FollowRepository.FindFollowByFollowerAndTarget(followerId, targetId)
}

func (u *FollowUsecase) FindFollowsByUserId(userId string) ([]*model.SimpleUser, error) {
	return u.FollowRepository.FindFollowsByUserId(userId)
}

func (u *FollowUsecase) FindFollowersByUserId(userId string) ([]*model.SimpleUser, error) {
	return u.FollowRepository.FindFollowersByUserId(userId)
}

func (u *FollowUsecase) FindNostrFollowPublicKeys(userId string) ([]string, error) {
	return u.FollowRepository.FindNostrFollowPublicKeys(userId)
}

func (u *FollowUsecase) CheckIsFollowing(followerId, targetId string) (bool, error) {
	return u.FollowRepository.CheckIsFollowing(followerId, targetId)
}

func (u *FollowUsecase) Delete(followerId, targetId string) error {
	return u.FollowRepository.Delete(followerId, targetId)
}

func (u *FollowUsecase) GetAllFollowingNostrPubKeys() ([]string, error) {
	return u.FollowRepository.GetAllFollowingNostrPubKeys()
}

func (u *FollowUsecase) GetAllLocalUserNostrPubKeys() ([]string, error) {
	return u.FollowRepository.GetAllLocalUserNostrPubKeys()
}

func (u *FollowUsecase) GetLatestNostrRemoteFollow() (*model.Follow, error) {
	return u.FollowRepository.GetLatestNostrRemoteFollow()
}
