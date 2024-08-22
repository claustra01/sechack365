package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
)

type UserRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *UserRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	row, err := repo.SqlHandler.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var user model.User
		if err = row.Scan(&user.Id, &user.UserId, &user.Host, &user.EncryptedPassword, &user.DisplayName, &user.Profile); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserRepository) FindById(id string) (*model.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.UserId, &user.Host, &user.EncryptedPassword, &user.DisplayName, &user.Profile); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (repo *UserRepository) FindByUserId(userId string) (*model.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.User)
		if err = row.Scan(&user.Id, &user.UserId, &user.Host, &user.EncryptedPassword, &user.DisplayName, &user.Profile); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

type ApUserRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *ApUserRepository) FindByUserId(userId string) (*model.ApUser, error) {
	row, err := repo.SqlHandler.Query(`
		SELECT users.id, users.user_id, host, encrypted_password, display_name, profile, public_key, private_key
		FROM users, ap_user_identifiers WHERE users.user_id = $1 AND users.id = ap_user_identifiers.user_id;
	`, userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var user = new(model.ApUser)
		if err = row.Scan(&user.Id, &user.UserId, &user.Host, &user.EncryptedPassword, &user.DisplayName, &user.Profile, &user.PublicKey, &user.PrivateKey); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}
