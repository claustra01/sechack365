package repository

import (
	"time"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type PostRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *PostRepository) Create(userId, content string) (*model.Post, error) {
	uuid := util.NewUuid()
	row, err := r.SqlHandler.Query(`
		INSERT INTO posts (id, user_id, content)
		VALUES ($1, $2, $3)
		RETURNING *;
	`, uuid, userId, content)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var post = new(model.Post)
		if err := row.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		return post, nil
	}
	return nil, nil
}

func (r *PostRepository) FindById(id string) (*model.PostWithUser, error) {
	row, err := r.SqlHandler.Query(`
		SELECT posts.*, users.username, users.host, users.protocol, users.display_name, users.profile, users.icon FROM posts JOIN users ON posts.user_id = users.id
		WHERE posts.id = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var post = new(model.PostWithUser)
		if err := row.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.User.Username, &post.User.Host, &post.User.Protocol, &post.User.DisplayName, &post.User.Profile, &post.User.Icon); err != nil {
			return nil, err
		}
		return post, nil
	}
	return nil, nil
}

func (r *PostRepository) FindTimeline(createdAt time.Time, limit int) ([]*model.PostWithUser, error) {
	row, err := r.SqlHandler.Query(`
		SELECT posts.*, users.username, users.host, users.protocol, users.display_name, users.profile, users.icon FROM posts JOIN users ON posts.user_id = users.id
		WHERE posts.created_at < $1
		ORDER BY posts.created_at DESC LIMIT $2 ;
	`, createdAt, limit)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var posts []*model.PostWithUser
	for row.Next() {
		var post = new(model.PostWithUser)
		if err := row.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.User.Username, &post.User.Host, &post.User.Protocol, &post.User.DisplayName, &post.User.Profile, &post.User.Icon); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) FindUserTimeline(userId string, createdAt time.Time, limit int) ([]*model.PostWithUser, error) {
	row, err := r.SqlHandler.Query(`
		SELECT posts.*, users.username, users.host, users.protocol, users.display_name, users.profile, users.icon FROM posts JOIN users ON posts.user_id = users.id
		WHERE user_id = $1 AND posts.created_at < $2
		ORDER BY posts.created_at DESC LIMIT $3;
	`, userId, createdAt, limit)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var posts []*model.PostWithUser
	for row.Next() {
		var post = new(model.PostWithUser)
		if err := row.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.User.Username, &post.User.Host, &post.User.Protocol, &post.User.DisplayName, &post.User.Profile, &post.User.Icon); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) Delete(id string) error {
	_, err := r.SqlHandler.Exec(`
		DELETE FROM posts WHERE id = $1;
	`, id)
	return err
}
