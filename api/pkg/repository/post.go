package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type PostRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *PostRepository) Create(userId, content string) error {
	uuid := util.NewUuid()
	query := `
		INSERT INTO posts (id, user_id, content)
		VALUES ($1, $2, $3)
		RETURNING *;
	`
	if _, err := r.SqlHandler.Exec(query, uuid, userId, content); err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) FindById(id string) (*model.PostWithUser, error) {
	post := new(model.PostWithUser)
	query := `
		SELECT posts.*,
		CASE
			WHEN users.protocol = 'local' THEN '@' || users.username
			WHEN users.protocol = 'activitypub' THEN ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
			WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
		END AS user_username,
		users.protocol AS user_protocol,
		users.display_name AS user_display_name,
		users.icon AS user_icon
		FROM posts
		JOIN users ON posts.user_id = users.id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE posts.id = $1;
	`
	err := r.SqlHandler.Get(post, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) FindTimeline(createdAt time.Time, limit int) ([]*model.PostWithUser, error) {
	rawPosts := make([]model.PostWithUser, 0)
	query := `
		SELECT posts.*,
		CASE
			WHEN users.protocol = 'local' THEN '@' || users.username
			WHEN users.protocol = 'activitypub' THEN ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
			WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
		END AS user_username,
		users.protocol AS user_protocol,
		users.display_name AS user_display_name,
		users.icon AS user_icon
		FROM posts
		JOIN users ON posts.user_id = users.id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE posts.created_at < $1
		ORDER BY posts.created_at DESC LIMIT $2;
	`
	if err := r.SqlHandler.Select(&rawPosts, query, createdAt, limit); err != nil {
		return nil, err
	}
	posts := make([]*model.PostWithUser, 0)
	for _, rawPost := range rawPosts {
		posts = append(posts, &rawPost)
	}
	return posts, nil
}

func (r *PostRepository) FindUserTimeline(userId string, createdAt time.Time, limit int) ([]*model.PostWithUser, error) {
	rawPosts := make([]model.PostWithUser, 0)
	query := `
		SELECT posts.*,
		CASE
			WHEN users.protocol = 'local' THEN '@' || users.username
			WHEN users.protocol = 'activitypub' THEN ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
			WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
		END AS user_username,
		users.protocol AS user_protocol,
		users.display_name AS user_display_name,
		users.icon AS user_icon
		FROM posts
		JOIN users ON posts.user_id = users.id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE user_id = $1 AND posts.created_at < $2
		ORDER BY posts.created_at DESC LIMIT $3;
	`
	if err := r.SqlHandler.Select(&rawPosts, query, userId, createdAt, limit); err != nil {
		return nil, err
	}
	posts := make([]*model.PostWithUser, 0)
	for _, rawPost := range rawPosts {
		posts = append(posts, &rawPost)
	}
	return posts, nil
}

func (r *PostRepository) Delete(id string) error {
	_, err := r.SqlHandler.Exec(`DELETE FROM posts WHERE id = $1;`, id)
	return err
}
