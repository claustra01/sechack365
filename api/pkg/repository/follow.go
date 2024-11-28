package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type FollowRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *FollowRepository) Create(followerId, followeeId string) (*model.Follow, error) {
	follow := new(model.Follow)
	uuid := util.NewUuid()
	query := `
		INSERT INTO follows (id, follower_id, followee_id)
		VALUES ($1, $2, $3)
		RETURNING *;
	`
	if err := r.SqlHandler.Get(follow, query, uuid, followerId, followeeId); err != nil {
		return nil, err
	}
	return follow, nil
}

func (r *FollowRepository) UpdateAcceptFollow(followerId, followeeId string) (*model.Follow, error) {
	follow := new(model.Follow)
	query := `
		UPDATE follows SET is_accepted = true
		WHERE follower_id = $1 AND followee_id = $2
		RETURNING *;
	`
	if err := r.SqlHandler.Get(follow, query, followerId, followeeId); err != nil {
		return nil, err
	}
	return follow, nil
}

func (r *FollowRepository) FindFollowsByUserId(userId string) ([]*model.User, error) {
	var users []*model.User
	if err := r.SqlHandler.Select(&users, `SELECT users.* FROM users JOIN follows ON users.id = follows.followee_id WHERE follows.follower_id = $1;`, userId); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *FollowRepository) FindFollowersByUserId(userId string) ([]*model.User, error) {
	var users []*model.User
	if err := r.SqlHandler.Select(&users, `SELECT users.* FROM users JOIN follows ON users.id = follows.follower_id WHERE follows.followee_id = $1;`, userId); err != nil {
		return nil, err
	}
	return users, nil
}
