package controller

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/repository"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type PostController struct {
	PostUsecase usecase.PostUsecase
}

func NewPostController(conn model.ISqlHandler) *PostController {
	return &PostController{
		PostUsecase: usecase.PostUsecase{
			PostRepository: &repository.PostRepository{
				SqlHandler: conn,
			},
		},
	}
}

func (c *PostController) Create(userId, content string) error {
	return c.PostUsecase.Create(userId, content)
}

func (c *PostController) FindById(id string) (*model.PostWithUser, error) {
	return c.PostUsecase.FindById(id)
}

func (c *PostController) FindTimeline(offset int, limit int) ([]*model.PostWithUser, error) {
	return c.PostUsecase.FindTimeline(offset, limit)
}

func (c *PostController) FindUserTimeline(userId string, offset int, limit int) ([]*model.PostWithUser, error) {
	return c.PostUsecase.FindUserTimeline(userId, offset, limit)
}

func (c *PostController) DeleteById(id string) error {
	return c.PostUsecase.DeleteById(id)
}

func (c *PostController) GetLatestNostrRemotePost() (*model.Post, error) {
	return c.PostUsecase.GetLatestNostrRemotePost()
}

func (c *PostController) InsertNostrRemotePosts(events []*model.NostrEvent) error {
	return c.PostUsecase.InsertNostrRemotePosts(events)
}
