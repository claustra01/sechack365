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
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.EncryptedPassword, &user.DisplayName, &user.Profile, &user.CreatedAt, &user.UpdatedAt); err != nil {
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
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.EncryptedPassword, &user.DisplayName, &user.Profile, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (repo *UserRepository) FindByUsername(username string) (*model.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE username = $1;", username)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.EncryptedPassword, &user.DisplayName, &user.Profile, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (repo *UserRepository) Insert(user *model.User) error {
	_, err := repo.SqlHandler.Execute(`
		INSERT INTO users (id, username, host, encrypted_password, display_name, profile)
			VALUES ($1, $2, $3, $4, $5, $6);
	`, user.Id, user.Username, user.Host, user.EncryptedPassword, user.DisplayName, user.Profile)
	return err
}

type ApUserIdentifierRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *ApUserIdentifierRepository) Insert(userId string) error {
	publicKey, privateKey, err := util.GenerateKeyPair()
	if err != nil {
		return err
	}
	_, err = repo.SqlHandler.Execute(`
		INSERT INTO ap_user_identifiers (user_id, public_key, private_key)
			VALUES ($1, $2, $3);
	`, userId, publicKey, privateKey)
	return err
}

type ApUserRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *ApUserRepository) FindByUsername(username string) (*model.ApUser, error) {
	row, err := repo.SqlHandler.Query(`
		SELECT users.id, users.username, host, encrypted_password, display_name, profile, public_key, private_key, users.created_at, users.updated_at
		FROM users, ap_user_identifiers WHERE users.username = $1 AND users.id = ap_user_identifiers.user_id;
	`, username)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.ApUser)
		if err = row.Scan(&user.Id, &user.Username, &user.Host, &user.EncryptedPassword, &user.DisplayName, &user.Profile, &user.PublicKey, &user.PrivateKey, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}
