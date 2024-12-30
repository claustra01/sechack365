package repository

import (
	"database/sql"
	"errors"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type UserRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *UserRepository) CreateLocalUser(username, password, displayName, profile, icon, host string) error {
	// create user record
	uuid := util.NewUuid()
	hashedPassword, err := util.GenerateHash(password)
	if err != nil {
		return err
	}
	query := `
		INSERT INTO users (id, username, protocol, hashed_password, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, username, model.ProtocolLocal, hashedPassword, displayName, profile, icon); err != nil {
		return err
	}

	// create ap_user_identifier record
	pubKey, prvKey, err := util.GenerateApKeyPair()
	if err != nil {
		return err
	}
	query = `
		INSERT INTO ap_user_identifiers (user_id, local_username, host, public_key, private_key)
		VALUES ($1, $2, $3, $4, $5);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, username, host, pubKey, prvKey); err != nil {
		return err
	}

	// create nostr_user_identifier record
	pubKey, prvKey, err = util.GenerateNostrKeyPair()
	if err != nil {
		return err
	}
	npub, err := util.EncodeNpub(pubKey)
	if err != nil {
		return err
	}
	nsec, err := util.EncodeNsec(prvKey)
	if err != nil {
		return err
	}
	query = `
		INSERT INTO nostr_user_identifiers (user_id, public_key, private_key)
		VALUES ($1, $2, $3);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, npub, nsec); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) CreateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) error {
	// create user record
	uuid := util.NewUuid()
	query := `
		INSERT INTO users (id, protocol, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, model.ProtocolActivityPub, user.DisplayName, user.Profile, user.Icon); err != nil {
		return err
	}

	// create ap_user_identifier record
	query = `
		INSERT INTO ap_user_identifiers (user_id, local_username, host)
		VALUES ($1, $2, $3);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, identifier.LocalUsername, identifier.Host); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) CreateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) error {
	// create user record
	uuid := util.NewUuid()
	query := `
		INSERT INTO users (id, protocol, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`
	if _, err := r.SqlHandler.Exec(query, uuid, model.ProtocolNostr, user.DisplayName, user.Profile, user.Icon); err != nil {
		return err
	}

	// create nostr_user_identifier record
	query = `
		INSERT INTO nostr_user_identifiers (user_id, public_key)
		VALUES ($1, $2)
		RETURNING user_id, public_key;
	`
	if _, err := r.SqlHandler.Exec(query, uuid, identifier.PublicKey); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindAll() ([]*model.UserWithIdentifiers, error) {
	var users []*model.UserWithIdentifiers
	query := `
		SELECT
			users.id,
			users.username,
			users.protocol,
			users.display_name,
			users.profile,
			users.icon,
			users.created_at,
			users.updated_at,
			COALESCE(ap_user_identifiers.local_username, '') AS "identifiers.activitypub.local_username",
			COALESCE(ap_user_identifiers.host, '') AS "identifiers.activitypub.host",
			COALESCE(ap_user_identifiers.public_key, '') AS "identifiers.activitypub.public_key",
			COALESCE(nostr_user_identifiers.public_key, '') AS "identifiers.nostr.public_key",
			(SELECT COUNT(*) FROM posts WHERE posts.user_id = users.id) AS post_count,
			(SELECT COUNT(*) FROM follows WHERE follows.follower_id = users.id) AS follow_count,
			(SELECT COUNT(*) FROM follows WHERE follows.target_id = users.id) AS follower_count
		FROM users
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id;
	`
	if err := r.SqlHandler.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindById(id string) (*model.UserWithIdentifiers, error) {
	user := new(model.UserWithIdentifiers)
	query := `
		SELECT
			users.id,
			users.username,
			users.protocol,
			users.display_name,
			users.profile,
			users.icon,
			users.created_at,
			users.updated_at,
			COALESCE(ap_user_identifiers.local_username, '') AS "identifiers.activitypub.local_username",
			COALESCE(ap_user_identifiers.host, '') AS "identifiers.activitypub.host",
			COALESCE(ap_user_identifiers.public_key, '') AS "identifiers.activitypub.public_key",
			COALESCE(nostr_user_identifiers.public_key, '') AS "identifiers.nostr.public_key",
			(SELECT COUNT(*) FROM posts WHERE posts.user_id = users.id) AS post_count,
			(SELECT COUNT(*) FROM follows WHERE follows.follower_id = users.id) AS follow_count,
			(SELECT COUNT(*) FROM follows WHERE follows.target_id = users.id) AS follower_count
		FROM users
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE users.id = $1;
	`
	if err := r.SqlHandler.Get(user, query, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByLocalUsername(username string) (*model.UserWithIdentifiers, error) {
	user := new(model.UserWithIdentifiers)
	query := `
		SELECT
			users.id,
			users.username,
			users.protocol,
			users.display_name,
			users.profile,
			users.icon,
			users.created_at,
			users.updated_at,
			COALESCE(ap_user_identifiers.local_username, '') AS "identifiers.activitypub.local_username",
			COALESCE(ap_user_identifiers.host, '') AS "identifiers.activitypub.host",
			COALESCE(ap_user_identifiers.public_key, '') AS "identifiers.activitypub.public_key",
			COALESCE(nostr_user_identifiers.public_key, '') AS "identifiers.nostr.public_key",
			(SELECT COUNT(*) FROM posts WHERE posts.user_id = users.id) AS post_count,
			(SELECT COUNT(*) FROM follows WHERE follows.follower_id = users.id) AS follow_count,
			(SELECT COUNT(*) FROM follows WHERE follows.target_id = users.id) AS follower_count
		FROM users
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE users.username = $1 AND users.protocol = $2;
	`
	err := r.SqlHandler.Get(user, query, username, model.ProtocolLocal)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByApUsername(username, host string) (*model.UserWithIdentifiers, error) {
	user := new(model.UserWithIdentifiers)
	query := `
		SELECT
			users.id,
			users.username,
			users.protocol,
			users.display_name,
			users.profile,
			users.icon,
			users.created_at,
			users.updated_at,
			ap_user_identifiers.local_username AS "identifiers.activitypub.local_username",
			ap_user_identifiers.host AS "identifiers.activitypub.host",
			ap_user_identifiers.public_key AS "identifiers.activitypub.public_key",
			(SELECT COUNT(*) FROM posts WHERE posts.user_id = users.id) AS post_count,
			(SELECT COUNT(*) FROM follows WHERE follows.follower_id = users.id) AS follow_count,
			(SELECT COUNT(*) FROM follows WHERE follows.target_id = users.id) AS follower_count
		FROM users
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		WHERE users.protocol = $1 AND ap_user_identifiers.local_username = $2 AND ap_user_identifiers.host = $3;
	`
	err := r.SqlHandler.Get(user, query, model.ProtocolActivityPub, username, host)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByNostrPublicKey(publicKey string) (*model.UserWithIdentifiers, error) {
	user := new(model.UserWithIdentifiers)
	query := `
		SELECT
			users.id,
			users.username,
			users.protocol,
			users.display_name,
			users.profile,
			users.icon,
			users.created_at,
			users.updated_at,
			nostr_user_identifiers.public_key AS "identifiers.nostr.public_key",
			(SELECT COUNT(*) FROM posts WHERE posts.user_id = users.id) AS post_count,
			(SELECT COUNT(*) FROM follows WHERE follows.follower_id = users.id) AS follow_count,
			(SELECT COUNT(*) FROM follows WHERE follows.target_id = users.id) AS follower_count
		FROM users
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE users.protocol = $1 AND nostr_user_identifiers.public_key = $2;
	`
	err := r.SqlHandler.Get(user, query, model.ProtocolNostr, publicKey)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) error {
	// get user id
	var userId string
	query := "SELECT user_id FROM ap_user_identifiers WHERE local_username = $1 AND host = $2;"
	if err := r.SqlHandler.Get(&userId, query, user.Username, identifier.Host); err != nil {
		return err
	}
	// update user record
	query = `
		UPDATE users
		SET display_name = $1, profile = $2, icon = $3
		WHERE id = $4;
	`
	if _, err := r.SqlHandler.Exec(query, user.DisplayName, user.Profile, user.Icon, userId); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) error {
	// get user id
	var userId string
	query := "SELECT user_id FROM nostr_user_identifiers WHERE public_key = $1;"
	if err := r.SqlHandler.Get(&userId, query, identifier.PublicKey); err != nil {
		return err
	}
	// update user record
	query = `
		UPDATE users
		SET display_name = $1, profile = $2, icon = $3
		WHERE id = $4;
	`
	if _, err := r.SqlHandler.Exec(query, user.DisplayName, user.Profile, user.Icon, userId); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteById(id string) error {
	if _, err := r.SqlHandler.Exec("DELETE FROM ap_user_identifiers WHERE user_id = $1;", id); err != nil {
		return err
	}
	if _, err := r.SqlHandler.Exec("DELETE FROM nostr_user_identifiers WHERE user_id = $1;", id); err != nil {
		return err
	}
	if _, err := r.SqlHandler.Exec("DELETE FROM users WHERE id = $1;", id); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindWithHashedPassword(username string) (*model.User, error) {
	user := new(model.User)
	// NOTE: should be a local user
	query := `
		SELECT * FROM users WHERE username = $1 AND protocol = $2;
	`
	if err := r.SqlHandler.Get(user, query, username, model.ProtocolLocal); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetNostrPrivKey(id string) (string, error) {
	var privKey string
	if err := r.SqlHandler.Get(&privKey, "SELECT private_key FROM nostr_user_identifiers WHERE user_id = $1;", id); err != nil {
		return "", err
	}
	return privKey, nil
}
