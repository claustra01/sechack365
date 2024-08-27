package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
)

type FollowRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *FollowRepository) Insert(follower string, followee string) (*model.Follow, error) {
	row, err := repo.SqlHandler.Query(`
		INSERT INTO follow (follower, followee)
		VALUES ($1, $2)
		RETURNING *;
	`, follower, followee)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		var follow = new(model.Follow)
		if err = row.Scan(&follow.Follower, &follow.Followee, &follow.IsAccepted, &follow.CreatedAt, &follow.UpdatedAt); err != nil {
			return nil, err
		}
		return follow, nil
	}
	return nil, nil
}
