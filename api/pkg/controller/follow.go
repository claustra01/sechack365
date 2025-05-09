package controller

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/repository"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type FollowController struct {
	FollowUsecase usecase.FollowUsecase
}

func NewFollowController(conn model.ISqlHandler) *FollowController {
	return &FollowController{
		FollowUsecase: usecase.FollowUsecase{
			FollowRepository: &repository.FollowRepository{
				SqlHandler: conn,
			},
		},
	}
}

func (c *FollowController) Create(followerId, targetId string) error {
	return c.FollowUsecase.Create(followerId, targetId)
}

func (c *FollowController) UpdateAcceptFollow(followerId, targetId string) error {
	return c.FollowUsecase.UpdateAcceptFollow(followerId, targetId)
}

func (c *FollowController) FindFollowByFollowerAndTarget(followerId, targetId string) (*model.Follow, error) {
	return c.FollowUsecase.FindFollowByFollowerAndTarget(followerId, targetId)
}

func (c *FollowController) FindFollowsByUserId(userId string) ([]*model.SimpleUser, error) {
	return c.FollowUsecase.FindFollowsByUserId(userId)
}

func (c *FollowController) FindFollowersByUserId(userId string) ([]*model.SimpleUser, error) {
	return c.FollowUsecase.FindFollowersByUserId(userId)
}

func (c *FollowController) FindActivityPubRemoteFollowers(userId string) ([]string, error) {
	return c.FollowUsecase.FindActivityPubRemoteFollowers(userId)
}

func (c *FollowController) FindNostrFollowPublicKeys(userId string) ([]string, error) {
	return c.FollowUsecase.FindNostrFollowPublicKeys(userId)
}

func (c *FollowController) CheckIsFollowing(followerId, targetId string) (bool, error) {
	return c.FollowUsecase.CheckIsFollowing(followerId, targetId)
}

func (c *FollowController) Delete(followerId, targetId string) error {
	return c.FollowUsecase.Delete(followerId, targetId)
}
