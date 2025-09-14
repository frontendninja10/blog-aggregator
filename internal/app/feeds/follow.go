package feeds

import (
	"context"
	"fmt"
	"time"

	"github.com/frontendninja10/blog-aggregator/internal/app"
	"github.com/frontendninja10/blog-aggregator/internal/database"
	"github.com/google/uuid"
)


func Follow(s *app.State, cmd app.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	ctx := context.Background()

	feed, err := s.DB.GetFeed(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

	createFeedFollow := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = s.DB.CreateFeedFollow(ctx, createFeedFollow)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	return nil
}