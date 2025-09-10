package main

import (
	"context"
	"fmt"
)

func getFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
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