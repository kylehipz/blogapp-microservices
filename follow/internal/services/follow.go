package services

import (
	"context"

	"github.com/google/uuid"
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
	follow, err := f.dbClient.FollowUser(ctx, followerId, followeeId)
	if err != nil {
		return nil, err
	}

	return follow, nil
}

func (f *FollowService) UnfollowUser(
	ctx context.Context,
	follower string,
	followee string,
) error {
	followerID, err := uuid.Parse(follower)
	if err != nil {
		return err
	}

	followeeID, err := uuid.Parse(follower)
	if err != nil {
		return err
	}

	err = f.Queries.UnfollowUser(ctx, db.UnfollowUserParams{
		Follower: followerID,
		Followee: followeeID,
	})
	if err != nil {
		return err
	}

	return nil
}
