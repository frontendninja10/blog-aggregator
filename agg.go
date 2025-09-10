package main

import (
	"context"
	"fmt"
)

func agg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	ctx := context.Background()


	feed, err := fetchFeed(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("error fetching feed")
	}

	fmt.Printf("Feed: %+v\n", feed)
	return nil
}