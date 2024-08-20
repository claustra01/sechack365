package repository

import "github.com/claustra01/sechack365/pkg/model"

type UserRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *UserRepository) FindAll() (users []*model.User, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var user model.User
		if err = row.Scan(&user.Id, &user.UserId, &user.Host, &user.DisplayName, &user.Profile); err != nil {
			return
		}
		users = append(users, &user)
	}
	return
}

func (repo *UserRepository) FindById(id string) (user *model.User, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		user = new(model.User)
		if err = row.Scan(&user.Id, &user.UserId, &user.Host, &user.DisplayName, &user.Profile); err != nil {
			return
		}
		return
	}
	return nil, nil
}

func (repo *UserRepository) FindByUserId(userId string) (user *model.User, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		user = new(model.User)
		if err = row.Scan(&user.Id, &user.UserId, &user.Host, &user.DisplayName, &user.Profile); err != nil {
			return
		}
		return
	}
	return nil, nil
}
