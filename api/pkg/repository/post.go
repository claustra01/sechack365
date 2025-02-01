package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type PostRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *PostRepository) Create(id, userId, content string) error {
	query := `
		INSERT INTO posts (id, protocol, user_id, content)
		VALUES ($1, $2, $3, $4);
	`
	if _, err := r.SqlHandler.Exec(query, id, model.ProtocolLocal, userId, content); err != nil {
		return cerror.Wrap(err, "failed to create post")
	}
	return nil
}

func (r *PostRepository) FindById(id string) (*model.PostWithUser, error) {
	post := new(model.PostWithUser)
	query := `
		SELECT posts.*,
		users.id AS "user.id",
		CASE
			WHEN users.protocol = 'local' THEN '@' || users.username
			WHEN users.protocol = 'activitypub' THEN '@' || ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
			WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
		END AS "user.username",
		users.protocol AS "user.protocol",
		users.display_name AS "user.display_name",
		users.icon AS "user.icon"
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
		return nil, cerror.Wrap(err, "failed to get post by id")
	}
	return post, nil
}

func (r *PostRepository) FindTimeline(offset int, limit int) ([]*model.PostWithUser, error) {
	var posts []*model.PostWithUser
	query := `
		SELECT posts.*,
		users.id AS "user.id",
		CASE
			WHEN users.protocol = 'local' THEN '@' || users.username
			WHEN users.protocol = 'activitypub' THEN '@' || ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
			WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
		END AS "user.username",
		users.protocol AS "user.protocol",
		users.display_name AS "user.display_name",
		users.icon AS "user.icon"
		FROM posts
		JOIN users ON posts.user_id = users.id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		ORDER BY posts.created_at DESC LIMIT $1 OFFSET $2;
	`
	if err := r.SqlHandler.Select(&posts, query, limit, offset); err != nil {
		return nil, cerror.Wrap(err, "failed to get timeline")
	}
	return posts, nil
}

func (r *PostRepository) FindUserTimeline(userId string, offset int, limit int) ([]*model.PostWithUser, error) {
	var posts []*model.PostWithUser
	query := `
		SELECT posts.*,
		users.id AS "user.id",
		CASE
			WHEN users.protocol = 'local' THEN '@' || users.username
			WHEN users.protocol = 'activitypub' THEN '@' || ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
			WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
		END AS "user.username",
		users.protocol AS "user.protocol",
		users.display_name AS "user.display_name",
		users.icon AS "user.icon"
		FROM posts
		JOIN users ON posts.user_id = users.id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE posts.user_id = $1
		ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3;
	`
	if err := r.SqlHandler.Select(&posts, query, userId, limit, offset); err != nil {
		return nil, cerror.Wrap(err, "failed to get user's timeline")
	}
	return posts, nil
}

func (r *PostRepository) DeleteById(id string) error {
	_, err := r.SqlHandler.Exec(`DELETE FROM posts WHERE id = $1;`, id)
	return cerror.Wrap(err, "failed to delete post")
}

func (r *PostRepository) InsertApRemotePost(userId string, note *model.ApNoteActivity) error {
	uuid := util.NewUuid()
	query := `
		INSERT INTO posts (id, protocol, user_id, content, created_at)
		VALUES ($1, $2, $3, $4, $5);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, model.ProtocolActivityPub, userId, note.Content, note.Published); err != nil {
		return cerror.Wrap(err, "failed to insert ap remote post")
	}
	return nil
}

func (r *PostRepository) GetLatestNostrRemotePost() (*model.Post, error) {
	post := new(model.Post)
	query := `
		SELECT * FROM posts
		WHERE protocol = $1
		ORDER BY posts.created_at DESC LIMIT 1;
	`
	err := r.SqlHandler.Get(post, query, model.ProtocolNostr)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, cerror.Wrap(err, "failed to get latest nostr remote post")
	}
	return post, nil
}

func (r *PostRepository) InsertNostrRemotePosts(events []*model.NostrEvent) error {
	// FIXME: resolve N+1
	for _, event := range events {
		// get local user id
		var userId string
		query := `
			SELECT users.id FROM users
			JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
			WHERE nostr_user_identifiers.public_key = $1;
		`
		err := r.SqlHandler.Get(&userId, query, event.Pubkey)
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		} else if err != nil {
			return cerror.Wrap(err, "failed to insert nostr remote posts")
		}

		// check duplication
		var count int
		query = `
			SELECT COUNT(*) FROM posts
			WHERE user_id = $1 AND content = $2 AND created_at = $3;
		`
		if err := r.SqlHandler.Get(&count, query, userId, event.Content, time.Unix(int64(event.CreatedAt), 0)); err != nil {
			return cerror.Wrap(err, "failed to insert nostr remote posts")
		}
		if count > 0 {
			continue
		}

		// insert post
		uuid := util.NewUuid()
		query = `
			INSERT INTO posts (id, protocol, user_id, content, created_at)
			VALUES ($1, $2, $3, $4, $5);
		`
		if _, err := r.SqlHandler.Exec(query, uuid, model.ProtocolNostr, userId, event.Content, time.Unix(int64(event.CreatedAt), 0)); err != nil {
			return cerror.Wrap(err, "failed to insert nostr remote posts")
		}
	}
	return nil
}
