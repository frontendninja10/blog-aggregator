package feeds

import (
	"context"
	"fmt"
	"strconv"

	"github.com/frontendninja10/blog-aggregator/internal/app"
	"github.com/frontendninja10/blog-aggregator/internal/database"
)

func Browse(s *app.State, cmd app.Command, user database.User) error {
	limit := 2

	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})

	if err != nil {
		return fmt.Errorf("failed to get posts: %w", err)
	}

	fmt.Printf("Found %d posts for %s\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("	%v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("======================================")
	}

	return nil
}