package feeds

import (
	"context"
	"fmt"

	"github.com/frontendninja10/blog-aggregator/internal/app"
	"github.com/frontendninja10/blog-aggregator/internal/database"
)

func Unfollow(s *app.State, cmd app.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	feed, err := s.DB.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

	deleteFeedFollowArg := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.DB.DeleteFeedFollow(context.Background(), deleteFeedFollowArg); err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Printf("%s unfollowed successfully\n", feed.Name)

	return nil
}