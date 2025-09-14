package feeds

import (
	"context"
	"fmt"

	"github.com/frontendninja10/blog-aggregator/internal/app"
)

func ListFeeds(s *app.State, cmd app.Command) error {
	ctx := context.Background()

	feeds, err := s.DB.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch feeds: %w", err)
	}

	for _, feed := range feeds {
		
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("User: %s\n", feed.UserName)
		fmt.Println("============================================")
	}

	return nil
}