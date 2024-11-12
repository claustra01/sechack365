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

func (c *FollowController) Create(followerId, followeeId string) (*model.Follow, error) {
	return c.FollowUsecase.Create(followerId, followeeId)
}

func (c *FollowController) UpdateAcceptFollow(followerId, followeeId string) (*model.Follow, error) {
	return c.FollowUsecase.UpdateAcceptFollow(followerId, followeeId)
}

func (c *FollowController) FindFollowsByUserId(userId string) ([]*model.User, error) {
	return c.FollowUsecase.FindFollowsByUserId(userId)
}

func (c *FollowController) FindFollowersByUserId(userId string) ([]*model.User, error) {
	return c.FollowUsecase.FindFollowersByUserId(userId)
}
