package repository

import (
	"database/sql"
	"errors"

	"github.com/claustra01/sechack365/pkg/cerror"
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
		return cerror.Wrap(err, "failed to create follow")
	}
	return nil
}

func (r *FollowRepository) UpdateAcceptFollow(followerId, targetId string) error {
	query := `
		UPDATE follows SET is_accepted = true
		WHERE follower_id = $1 AND target_id = $2;
	`
	if _, err := r.SqlHandler.Exec(query, followerId, targetId); err != nil {
		return cerror.Wrap(err, "failed to accept follow")
	}
	return nil
}

func (r *FollowRepository) FindFollowByFollowerAndTarget(followerId, targetId string) (*model.Follow, error) {
	var follow model.Follow
	if err := r.SqlHandler.Get(&follow, "SELECT * FROM follows WHERE follower_id = $1 AND target_id = $2;", followerId, targetId); err != nil {
		return nil, cerror.Wrap(err, "failed to get follow")
	}
	return &follow, nil
}

func (r *FollowRepository) FindFollowsByUserId(userId string) ([]*model.SimpleUser, error) {
	var users []*model.SimpleUser
	query := `
		SELECT
			users.id,
			CASE
				WHEN users.protocol = 'local' THEN '@' || users.username
				WHEN users.protocol = 'activitypub' THEN '@' || ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
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
		return nil, cerror.Wrap(err, "failed to get follows")
	}
	return users, nil
}

func (r *FollowRepository) FindFollowersByUserId(userId string) ([]*model.SimpleUser, error) {
	var users []*model.SimpleUser
	query := `
		SELECT
			users.id,
			CASE
				WHEN users.protocol = 'local' THEN '@' || users.username
				WHEN users.protocol = 'activitypub' THEN '@' || ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
				WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
			END AS username,
			users.protocol,
			users.display_name,
			users.icon
		FROM users
		JOIN follows ON users.id = follows.follower_id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE follows.target_id = $1;
	`
	if err := r.SqlHandler.Select(&users, query, userId); err != nil {
		return nil, cerror.Wrap(err, "failed to get followers")
	}
	return users, nil
}

func (r *FollowRepository) FindActivityPubRemoteFollowers(userId string) ([]string, error) {
	var remoteFollowers []string
	query := `
		SELECT ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
		FROM users
		JOIN follows ON users.id = follows.follower_id
		JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		WHERE follows.target_id = $1 AND users.protocol = $2;
	`
	if err := r.SqlHandler.Select(&remoteFollowers, query, userId, model.ProtocolActivityPub); err != nil {
		return nil, cerror.Wrap(err, "failed to get activitypub remote followers")
	}
	return remoteFollowers, nil
}

func (r *FollowRepository) FindNostrFollowPublicKeys(userId string) ([]string, error) {
	var publicKeys []string
	query := `
		SELECT nostr_user_identifiers.public_key
		FROM nostr_user_identifiers
		JOIN follows ON nostr_user_identifiers.user_id = follows.target_id
		WHERE follows.follower_id = $1;
	`
	if err := r.SqlHandler.Select(&publicKeys, query, userId); err != nil {
		return nil, cerror.Wrap(err, "failed to get follows by public keys")
	}
	if len(publicKeys) == 0 {
		return []string{}, nil
	}
	return publicKeys, nil
}

func (r *FollowRepository) CheckIsFollowing(followerId, targetId string) (bool, error) {
	var followed bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM follows
			WHERE follower_id = $1 AND target_id = $2
		);
	`
	if err := r.SqlHandler.Get(&followed, query, followerId, targetId); err != nil {
		return false, cerror.Wrap(err, "failed to get if following")
	}
	return followed, nil
}

func (r *FollowRepository) Delete(followerId, targetId string) error {
	query := `
		DELETE FROM follows
		WHERE follower_id = $1 AND target_id = $2;
	`
	_, err := r.SqlHandler.Exec(query, followerId, targetId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return cerror.Wrap(err, "failed to delete follow")
	}
	return nil
}
