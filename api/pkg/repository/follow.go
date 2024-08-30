package repository

import (
	"github.com/claustra01/sechack365/pkg/model"
)

type FollowRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *FollowRepository) Insert(followerId string, followeeId string) (*model.Follow, error) {
	row, err := repo.SqlHandler.Query(`
		INSERT INTO follow (follower, followee)
		VALUES ($1, $2)
		RETURNING *;
	`, followerId, followeeId)
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

func (repo *FollowRepository) UpdateAcceptFollow(followerId string, followeeId string) (*model.Follow, error) {
	row, err := repo.SqlHandler.Query(`
		UPDATE follows SET is_accepted = true, updated_at = NOW()
		WHERE follower_id = $1 AND WHERE followee_id = $2
		RETURNING *;
	`, followerId, followeeId)
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
