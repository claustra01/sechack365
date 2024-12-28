package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type FollowRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *FollowRepository) Create(followerId, targetId string) error {
	uuid := util.NewUuid()
	query := `
		INSERT INTO follows (id, follower_id, target_id)
		VALUES ($1, $2, $3);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, followerId, targetId); err != nil {
		return err
	}
	return nil
}

func (r *FollowRepository) UpdateAcceptFollow(followerId, targetId string) error {
	query := `
		UPDATE follows SET is_accepted = true
		WHERE follower_id = $1 AND target_id = $2;
	`
	if _, err := r.SqlHandler.Exec(query, followerId, targetId); err != nil {
		return err
	}
	return nil
}

func (r *FollowRepository) FindFollowsByUserId(userId string) ([]*model.SimpleUser, error) {
	var users []*model.SimpleUser
	query := `
		SELECT
			CASE
				WHEN users.protocol = 'local' THEN '@' || users.username
				WHEN users.protocol = 'activitypub' THEN '@' ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
				WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
			END AS username,
			users.protocol,
			users.display_name,
			users.icon
		FROM users
		JOIN follows ON users.id = follows.target_id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE follows.follower_id = $1;
	`
	if err := r.SqlHandler.Select(&users, query, userId); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *FollowRepository) FindFollowersByUserId(userId string) ([]*model.SimpleUser, error) {
	var users []*model.SimpleUser
	query := `
		SELECT
			CASE
				WHEN users.protocol = 'local' THEN '@' || users.username
				WHEN users.protocol = 'activitypub' THEN '@' ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
				WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
			END AS username,
			users.protocol,
			users.display_name,
			users.icon
		FROM users
		JOIN follows ON users.id = follows.target_id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE follows.target_id = $1;
	`
	if err := r.SqlHandler.Select(&users, query, userId); err != nil {
		return nil, err
	}
	return users, nil
}
