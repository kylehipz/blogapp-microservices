package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
)

type FollowService struct {
	Queries *db.Queries
}

func (f *FollowService) FollowUser(
	ctx context.Context,
	follower string,
	followee string,
) (*db.Follow, error) {
	followerID, err := uuid.Parse(follower)
	if err != nil {
		return nil, err
	}

	followeeID, err := uuid.Parse(follower)
	if err != nil {
		return nil, err
	}

	follow, err := f.Queries.FollowUser(ctx, db.FollowUserParams{
		Follower: followerID,
		Followee: followeeID,
	})
	if err != nil {
		return nil, err
	}

	return &follow, nil
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
