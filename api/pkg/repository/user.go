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

type ApUserIdentifierRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *UserRepository) Create(username, host, protocol, password, displayName, profile, icon string) (*model.User, error) {
	user := new(model.User)
	uuid := util.NewUuid()
	hashedPassword, err := util.GenerateHash(password)
	if err != nil {
		return nil, err
	}
	query := `
		INSERT INTO users (id, username, host, protocol, hashed_password, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(user, query, uuid, username, host, protocol, hashedPassword, displayName, profile, icon); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.SqlHandler.Select(&users, "SELECT * FROM users;"); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindById(id string) (*model.User, error) {
	user := new(model.User)
	err := r.SqlHandler.Get(user, "SELECT * FROM users WHERE id = $1;", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByUsername(username, host string) (*model.User, error) {
	user := new(model.User)
	err := r.SqlHandler.Get(user, "SELECT * FROM users WHERE username = $1 AND host = $2;", username, host)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteById(id string) error {
	_, err := r.SqlHandler.Exec("DELETE FROM users WHERE id = $1;", id)
	return err
}

func (r *UserRepository) CreateRemoteUser(username, host, protocol, displayName, profile, icon string) (*model.User, error) {
	user := new(model.User)
	uuid := util.NewUuid()
	query := `
		INSERT INTO users (id, username, host, protocol, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(user, query, uuid, username, host, protocol, displayName, profile, icon); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateRemoteUser(username, host, displayName, profile, icon string) (*model.User, error) {
	user := new(model.User)
	query := `
		UPDATE users
		SET display_name = $1, profile = $2, icon = $3
		WHERE username = $4 AND host = $5
		RETURNING *;
	`
	if err := r.SqlHandler.Get(user, query, displayName, profile, icon, username, host); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *ApUserIdentifierRepository) Create(id string) (*model.ApUserIdentifier, error) {
	apUserIdentifier := new(model.ApUserIdentifier)
	pubKey, prvKey, err := util.GenerateApKeyPair()
	if err != nil {
		return nil, err
	}
	query := `
		INSERT INTO ap_user_identifiers (user_id, public_key, private_key)
		VALUES ($1, $2, $3)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(apUserIdentifier, query, id, pubKey, prvKey); err != nil {
		return nil, err
	}
	return apUserIdentifier, nil
}

func (r *ApUserIdentifierRepository) FindById(id string) (*model.ApUserIdentifier, error) {
	apUserIdentifier := new(model.ApUserIdentifier)
	err := r.SqlHandler.Get(apUserIdentifier, "SELECT * FROM ap_user_identifiers WHERE user_id = $1;", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return apUserIdentifier, nil
}

func (r *ApUserIdentifierRepository) DeleteById(id string) error {
	_, err := r.SqlHandler.Exec("DELETE FROM ap_user_identifiers WHERE user_id = $1;", id)
	return err
}
