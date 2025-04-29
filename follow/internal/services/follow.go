package services

import (
	"context"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type FollowService struct {
	Queries  *db.Queries
	dbClient db.DatabaseClient
}

func (f *FollowService) FollowUser(
	ctx context.Context,
	followerId string,
	followeeId string,
) (*types.Follow, error) {
	follow, err := f.dbClient.CreateFollow(ctx, followerId, followeeId)
	if err != nil {
		return nil, err
	}

	return follow, nil
}

func (f *FollowService) UnfollowUser(
	ctx context.Context,
	followerId string,
	followeeId string,
) error {
	if err := f.dbClient.DeleteFollow(ctx, followerId, followeeId); err != nil {
		return err
	}

	return nil
}
