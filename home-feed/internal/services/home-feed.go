package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
)

type HomeFeedService struct {
	Queries *db.Queries
}

func (h *HomeFeedService) GetHomeFeed(
	ctx context.Context,
	user string,
	createdAt string,
	limit int32,
) ([]db.Blog, error) {
	userID, err := uuid.Parse(user)
	if err != nil {
		return nil, err
	}

	// convert timestamp to time.Time
	t, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		panic(err)
	}

	blogs, err := h.Queries.GetHomeFeed(ctx, db.GetHomeFeedParams{
		Follower:  userID,
		CreatedAt: t,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}

	return blogs, nil
}
