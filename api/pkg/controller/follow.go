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

func (controller *FollowController) Insert(follower string, followee string) (*model.Follow, error) {
	return controller.FollowUsecase.Insert(follower, followee)
}
