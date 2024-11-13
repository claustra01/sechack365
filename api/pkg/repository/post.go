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

func (r *PostRepository) FindById(id string) (*model.Post, error) {
	row, err := r.SqlHandler.Query(`
		SELECT * FROM posts
		WHERE id = $1;
	`, id)
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

func (r *PostRepository) FindTimeline(createdAt time.Time, limit int) ([]*model.Post, error) {
	row, err := r.SqlHandler.Query(`
		SELECT * FROM posts
		WHERE created_at < $1
		LIMIT $2
	`, createdAt, limit)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var posts []*model.Post
	for row.Next() {
		var post = new(model.Post)
		if err := row.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) FindUserTimeline(userId string, createdAt time.Time, limit int) ([]*model.Post, error) {
	row, err := r.SqlHandler.Query(`
		SELECT * FROM posts
		WHERE user_id = $1 AND created_at < $2
		LIMIT $3
	`, userId, createdAt, limit)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var posts []*model.Post
	for row.Next() {
		var post = new(model.Post)
		if err := row.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
