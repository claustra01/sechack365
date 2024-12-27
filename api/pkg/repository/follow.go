package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type FollowRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *FollowRepository) Create(followerId, targetId string) error {
	uuid := util.NewUuid()
	query := `
		INSERT INTO follows (id, follower_id, target_id)
		VALUES ($1, $2, $3);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, followerId, targetId); err != nil {
		return err
	}
	return nil
}

func (r *FollowRepository) UpdateAcceptFollow(followerId, targetId string) error {
	query := `
		UPDATE follows SET is_accepted = true
		WHERE follower_id = $1 AND target_id = $2;
	`
	if _, err := r.SqlHandler.Exec(query, followerId, targetId); err != nil {
		return err
	}
	return nil
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
