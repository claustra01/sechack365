package usecase

import (
	"github.com/claustra01/sechack365/pkg/model"
)

type IPostRepository interface {
	Create(id, userId, content string) error
	FindById(id string) (*model.PostWithUser, error)
	FindTimeline(offset int, limit int) ([]*model.PostWithUser, error)
	FindUserTimeline(userId string, offset int, limit int) ([]*model.PostWithUser, error)
	DeleteById(id string) error
	// ActivityPub用
	InsertApRemotePost(userId string, note *model.ApNoteActivity) error
	// Nostr用
	GetLatestNostrRemotePost() (*model.Post, error)
	InsertNostrRemotePosts(events []*model.NostrEvent) error
}

type PostUsecase struct {
	PostRepository IPostRepository
}

func (u *PostUsecase) Create(id, userId, content string) error {
	return u.PostRepository.Create(id, userId, content)
}

func (u *PostUsecase) FindById(id string) (*model.PostWithUser, error) {
	return u.PostRepository.FindById(id)
}

func (u *PostUsecase) FindTimeline(offset int, limit int) ([]*model.PostWithUser, error) {
	return u.PostRepository.FindTimeline(offset, limit)
}

func (u *PostUsecase) FindUserTimeline(userId string, offset int, limit int) ([]*model.PostWithUser, error) {
	return u.PostRepository.FindUserTimeline(userId, offset, limit)
}

func (u *PostUsecase) DeleteById(id string) error {
	return u.PostRepository.DeleteById(id)
}

func (u *PostUsecase) InsertApRemotePost(userId string, note *model.ApNoteActivity) error {
	return u.PostRepository.InsertApRemotePost(userId, note)
}

func (u *PostUsecase) GetLatestNostrRemotePost() (*model.Post, error) {
	return u.PostRepository.GetLatestNostrRemotePost()
}

func (u *PostUsecase) InsertNostrRemotePosts(events []*model.NostrEvent) error {
	return u.PostRepository.InsertNostrRemotePosts(events)
}
