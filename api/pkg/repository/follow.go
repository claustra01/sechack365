package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type FollowRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *FollowRepository) Create(followerId, followeeId string) (*model.Follow, error) {
	uuid := util.NewUuid()
	row, err := r.SqlHandler.Query(`
		INSERT INTO follows (id, follower_id, followee_id)
		VALUES ($1, $2, $3)
		RETURNING *;
	`, uuid, followerId, followeeId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var follow = new(model.Follow)
		if err := row.Scan(&follow.Id, &follow.FollowerId, &follow.FolloweeId, &follow.IsAccepted, &follow.CreatedAt, &follow.UpdatedAt); err != nil {
			return nil, err
		}
		return follow, nil
	}
	return nil, nil
}

func (r *FollowRepository) UpdateAcceptFollow(followerId, followeeId string) (*model.Follow, error) {
	row, err := r.SqlHandler.Query(`
		UPDATE follows SET is_accepted = true
		WHERE follower_id = $1 AND followee_id = $2
		RETURNING *;
	`, followerId, followeeId)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var follow = new(model.Follow)
		if err := row.Scan(&follow.Id, &follow.FollowerId, &follow.FolloweeId, &follow.IsAccepted, &follow.CreatedAt, &follow.UpdatedAt); err != nil {
			return nil, err
		}
		return follow, nil
	}
	return nil, nil
}

func (r *FollowRepository) FindFollowsByUserId(userId string) ([]*model.User, error) {
	rows, err := r.SqlHandler.Query(`
		SELECT users.* FROM users JOIN follows ON users.id = follows.follower_id WHERE follows.follower_id = $1;
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*model.User
	for rows.Next() {
		var user = new(model.User)
		if err := rows.Scan(&user.Id, &user.Username, &user.Host, &user.Protocol, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *FollowRepository) FindFollowersByUserId(userId string) ([]*model.User, error) {
	rows, err := r.SqlHandler.Query(`
		SELECT users.* FROM users JOIN follows ON users.id = follows.followee_id WHERE follows.followee_id = $1;
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*model.User
	for rows.Next() {
		var user = new(model.User)
		if err := rows.Scan(&user.Id, &user.Username, &user.Host, &user.Protocol, &user.DisplayName, &user.Profile, &user.Icon, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
