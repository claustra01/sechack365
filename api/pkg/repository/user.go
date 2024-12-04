package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type UserRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *UserRepository) CreateLocalUser(username, password, displayName, profile, icon, host string) (*model.UserWithIdentifiers, error) {
	newUser := new(model.UserWithIdentifiers)
	// create user record
	uuid := util.NewUuid()
	hashedPassword, err := util.GenerateHash(password)
	if err != nil {
		return nil, err
	}
	query := `
		INSERT INTO users (id, username, protocol, hashed_password, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(newUser, query, uuid, username, model.ProtocolLocal, hashedPassword, displayName, profile, icon); err != nil {
		return nil, err
	}

	// create ap_user_identifier record
	apUserIdentifier := new(model.ApUserIdentifier)
	pubKey, prvKey, err := util.GenerateApKeyPair()
	if err != nil {
		return nil, err
	}
	query = `
		INSERT INTO ap_user_identifiers (user_id, local_username, host, public_key, private_key)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(apUserIdentifier, query, newUser.Id, username, host, pubKey, prvKey); err != nil {
		return nil, err
	}

	// create nostr_user_identifier record
	nostrUserIdentifier := new(model.NostrUserIdentifier)
	pubKey, prvKey, err = util.GenerateNostrKeyPair()
	if err != nil {
		return nil, err
	}
	npub, err := util.EncodeNpub(pubKey)
	if err != nil {
		return nil, err
	}
	nsec, err := util.EncodeNsec(prvKey)
	if err != nil {
		return nil, err
	}
	query = `
		INSERT INTO nostr_user_identifiers (user_id, public_key, private_key)
		VALUES ($1, $2, $3)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(nostrUserIdentifier, query, newUser.Id, npub, nsec); err != nil {
		return nil, err
	}

	// return
	newUser.Identifiers = &model.Identifiers{
		Activitypub: apUserIdentifier,
		Nostr:       nostrUserIdentifier,
	}
	return newUser, nil
}

func (r *UserRepository) CreateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) (*model.UserWithIdentifiers, error) {
	newUser := new(model.UserWithIdentifiers)
	// create user record
	id := util.NewUuid()
	query := `
		INSERT INTO users (id, protocol, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(newUser, query, id, model.ProtocolActivityPub, user.DisplayName, user.Profile, user.Icon); err != nil {
		return nil, err
	}

	// create ap_user_identifier record
	query = `
		INSERT INTO ap_user_identifiers (user_id, local_username, host)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(identifier, query, newUser.Id, identifier.LocalUsername, identifier.Host); err != nil {
		return nil, err
	}

	// return
	newUser.Identifiers = &model.Identifiers{
		Activitypub: identifier,
	}
	return newUser, nil
}

func (r *UserRepository) CreateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) (*model.UserWithIdentifiers, error) {
	newUser := new(model.UserWithIdentifiers)
	// create user record
	id := util.NewUuid()
	query := `
		INSERT INTO users (id, protocol, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(newUser, query, id, model.ProtocolNostr, user.DisplayName, user.Profile, user.Icon); err != nil {
		return nil, err
	}

	// create nostr_user_identifier record
	query = `
		INSERT INTO nostr_user_identifiers (user_id, public_key)
		VALUES ($1, $2)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(identifier, query, newUser.Id, identifier.PublicKey, identifier.PrivateKey); err != nil {
		return nil, err
	}

	// return
	newUser.Identifiers = &model.Identifiers{
		Nostr: identifier,
	}
	return newUser, nil
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
			json_build_object(
				'activitypub', json_build_object(
					'local_username', ap_user_identifiers.local_username,
					'host', ap_user_identifiers.host,
					'public_key', ap_user_identifiers.public_key
				),
				'nostr', json_build_object(
					'public_key', nostr_user_identifiers.public_key
				)
			) AS identifiers
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
			json_build_object(
				'activitypub', json_build_object(
					'local_username', ap_user_identifiers.local_username,
					'host', ap_user_identifiers.host,
					'public_key', ap_user_identifiers.public_key
				),
				'nostr', json_build_object(
					'public_key', nostr_user_identifiers.public_key
				)
			) AS identifiers
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
			json_build_object(
				'activitypub', json_build_object(
					'local_username', ap_user_identifiers.local_username,
					'host', ap_user_identifiers.host,
					'public_key', ap_user_identifiers.public_key
				),
				'nostr', json_build_object(
					'public_key', nostr_user_identifiers.public_key
				)
			) AS identifiers
		FROM users
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE users.username = $1 AND users.protocol = $2;
	`
	if err := r.SqlHandler.Get(user, query, username, model.ProtocolLocal); err != nil {
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
			json_build_object(
				'activitypub', json_build_object(
					'local_username', ap_user_identifiers.local_username,
					'host', ap_user_identifiers.host,
					'public_key', ap_user_identifiers.public_key
				)
			) AS identifiers
		FROM users
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		WHERE users.protocol = $1 AND ap_user_identifiers.username = $2 AND ap_user_identifiers.host = $3;
	`
	if err := r.SqlHandler.Get(user, query, model.ProtocolActivityPub, username, host); err != nil {
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
			json_build_object(
				'nostr', json_build_object(
					'public_key', nostr_user_identifiers.public_key
				)
			) AS identifiers
		FROM users
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE users.protocol = $1 AND nostr_user_identifiers.public_key = $2;
	`
	if err := r.SqlHandler.Get(user, query, model.ProtocolNostr, publicKey); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) (*model.UserWithIdentifiers, error) {
	newUser := new(model.UserWithIdentifiers)
	// get user id
	query := `
		SELECT user_id FROM ap_user_identifiers WHERE local_username = $1 AND host = $2;
	`
	var userId string
	if err := r.SqlHandler.Get(&userId, query, user.Username, identifier.Host); err != nil {
		return nil, err
	}
	// update user record
	query = `
		UPDATE users
		SET display_name = $1, profile = $2, icon = $3
		WHERE id = $4
		RETURNING *;
	`
	if err := r.SqlHandler.Get(newUser, query, user.DisplayName, user.Profile, user.Icon, userId); err != nil {
		return nil, err
	}
	newUser.Identifiers = &model.Identifiers{
		Activitypub: identifier,
	}
	return newUser, nil
}

func (r *UserRepository) UpdateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) (*model.UserWithIdentifiers, error) {
	newUser := new(model.UserWithIdentifiers)
	// get user id
	query := `
		SELECT user_id FROM nostr_user_identifiers WHERE public_key = $1;
	`
	var userId string
	if err := r.SqlHandler.Get(&userId, query, identifier.PublicKey); err != nil {
		return nil, err
	}
	// update user record
	query = `
		UPDATE users
		SET display_name = $1, profile = $2, icon = $3
		WHERE id = $4
		RETURNING *;
	`
	if err := r.SqlHandler.Get(newUser, query, user.DisplayName, user.Profile, user.Icon, userId); err != nil {
		return nil, err
	}
	newUser.Identifiers = &model.Identifiers{
		Nostr: identifier,
	}
	return newUser, nil
}

func (r *UserRepository) DeleteById(id string) error {
	if _, err := r.SqlHandler.Exec("DELETE FROM ap_user_identifiers WHERE user_id = $1;", id); err != nil {
		return err
	}
	if _, err := r.SqlHandler.Exec("DELETE FROM nostr_user_identifiers WHERE user_id = $1;", id); err != nil {
		return err
	}
	if _, err := r.SqlHandler.Exec("DELETE FROM users WHERE user_id = $1;", id); err != nil {
		return err
	}
	return nil
}
