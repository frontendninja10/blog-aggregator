package feeds

import (
	"context"
	"fmt"

	"github.com/frontendninja10/blog-aggregator/internal/app"
	"github.com/frontendninja10/blog-aggregator/internal/database"
)

func Following(s *app.State, cmd app.Command, user database.User) error {
	ctx := context.Background()

	feedFollows, err := s.DB.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follow for user: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("%s\n", feedFollow.FeedName)
	}
	return nil
}