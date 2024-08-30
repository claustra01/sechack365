package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type UserRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *UserRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	row, err := repo.SqlHandler.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var user model.User
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.Protocol, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserRepository) FindById(id string) (*model.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE id = $1;", id)
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

func (repo *UserRepository) FindByUsername(username string, host string) (*model.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE username = $1 AND host = $2;", username, host)
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

func (repo *UserRepository) Insert(username string, password string, host string, protocol string, display_name string, profile string, icon string) (*model.User, error) {
	uuid := util.NewUuid()
	var hashedPassword string
	if password != "" {
		var err error
		hashedPassword, err = util.GenerateHash(password)
		if err != nil {
			return nil, err
		}
	}
	row, err := repo.SqlHandler.Query(`
		INSERT INTO users (id, username, host, protocol, hashed_password, display_name, profile, icon)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, username, host, protocol, hashed_password, display_name, profile, icon, created_at, updated_at;
	`, uuid, username, host, protocol, hashedPassword, display_name, profile, icon)
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

func (repo *UserRepository) UpdateRemoteUser(username string, host string, display_name string, profile string, icon string) (*model.User, error) {
	row, err := repo.SqlHandler.Query(`
		UPDATE users SET display_name = $1, profile = $2, icon = $3, updated_at = NOW()
		WHERE username = $4 AND host = $5
		RETURNING id, username, host, hashed_password, display_name, profile, icon, created_at, updated_at;
	`, display_name, profile, icon, username, host)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

type ApUserIdentifierRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *ApUserIdentifierRepository) Insert(userId string, baseUrl string, inbox string, outbox string, publicKey string) (*model.ApUserIdentifier, error) {
	var privateKey string
	if publicKey == "" {
		var err error
		publicKey, privateKey, err = util.GenerateKeyPair()
		if err != nil {
			return nil, err
		}
	}
	row, err := repo.SqlHandler.Query(`
		INSERT INTO ap_user_identifiers (user_id, base_url, inbox, outbox, public_key, private_key)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING user_id, inbox, outbox, public_key, private_key;
	`, userId, baseUrl, inbox, outbox, publicKey, privateKey)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var apUserIdentifier = new(model.ApUserIdentifier)
		if err = row.Scan(&apUserIdentifier.UserId, &apUserIdentifier.Inbox, &apUserIdentifier.Outbox, &apUserIdentifier.PublicKey, &apUserIdentifier.PrivateKey); err != nil {
			return nil, err
		}
		return apUserIdentifier, nil
	}
	return nil, nil
}

type ApUserRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *ApUserRepository) FindById(id string) (*model.ApUser, error) {
	row, err := repo.SqlHandler.Query(`
		SELECT users.id, users.username, host, hashed_password, display_name, profile, icon, base_url, inbox, outbox, public_key, private_key, users.created_at, users.updated_at
		FROM users, ap_user_identifiers WHERE users.id = $1 AND users.id = ap_user_identifiers.user_id;
	`, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.ApUser)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.BaseUrl, &user.Inbox, &user.Outbox, &user.PublicKey, &user.PrivateKey, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (repo *ApUserRepository) FindByUsername(username string, host string) (*model.ApUser, error) {
	row, err := repo.SqlHandler.Query(`
		SELECT users.id, users.username, host, hashed_password, display_name, profile, icon, base_url, inbox, outbox, public_key, private_key, users.created_at, users.updated_at
		FROM users, ap_user_identifiers WHERE users.username = $1 AND users.host = $2 AND users.id = ap_user_identifiers.user_id;
	`, username, host)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.ApUser)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.HashedPassword, &user.DisplayName, &user.Profile, &user.Icon, &user.BaseUrl, &user.Inbox, &user.Outbox, &user.PublicKey, &user.PrivateKey, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}
