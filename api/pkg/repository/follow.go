package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type FollowRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *FollowRepository) Insert(follower string, followee string) (*model.Follow, error) {
	uuid := util.NewUuid()
	row, err := repo.SqlHandler.Query(`
		INSERT INTO follow (id, follower, followee)
		VALUES ($1, $2, $3)
		RETURNING *;
	`, uuid, follower, followee)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var follow = new(model.Follow)
		if err = row.Scan(&follow.Id, &follow.Follower, &follow.Followee, &follow.CreatedAt, &follow.UpdatedAt); err != nil {
			return nil, err
		}
		return follow, nil
	}
	return nil, nil
}
