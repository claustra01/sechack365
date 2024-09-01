package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type UserRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	row, err := r.SqlHandler.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var users []*model.User
	for row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.Protocol, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) FindById(id string) (*model.User, error) {
	row, err := r.SqlHandler.Query(`
		SELECT * FROM users
		WHERE id = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.Protocol, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (r *UserRepository) FindByUsername(username, host string) (*model.User, error) {
	row, err := r.SqlHandler.Query(`
		SELECT * FROM users
		WHERE username = $1 AND host = $2;
	`, username, host)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.Protocol, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (r *UserRepository) CreateRemoteUser(username, host, protocol, displayName, profile, icon string) (*model.User, error) {
	uuid := util.NewUuid()
	row, err := r.SqlHandler.Query(`
		INSERT INTO users (id, username, host, protocol, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *;
	`, uuid, username, host, protocol, displayName, profile, icon)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.Protocol, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (r *UserRepository) UpdateRemoteUser(username, host, displayName, profile, icon string) (*model.User, error) {
	row, err := r.SqlHandler.Query(`
		UPDATE users
		SET display_name = $1, profile = $2, icon = $3
		WHERE username = $4 AND host = $5
		RETURNING *;
	`, displayName, profile, icon, username, host)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.Protocol, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

type ApUserIdentifierRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *ApUserIdentifierRepository) FindById(id string) (*model.ApUserIdentifier, error) {
	row, err := r.SqlHandler.Query(`
		SELECT * FROM ap_user_identifiers
		WHERE user_id = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var apUserIdentifier = new(model.ApUserIdentifier)
		if err = row.Scan(&apUserIdentifier.UserId, &apUserIdentifier.PublicKey, &apUserIdentifier.PrivateKey, &apUserIdentifier.CreatedAt, &apUserIdentifier.UpdatedAt); err != nil {
			return nil, err
		}
		return apUserIdentifier, nil
	}
	return nil, nil
}
