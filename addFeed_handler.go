package main

import (
	"context"
	"fmt"
	"time"

	"github.com/frontendninja10/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func addFeedHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.name)
	}

	ctx := context.Background()

	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, feedParams)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	createFeedFollow := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = s.db.CreateFeedFollow(ctx, createFeedFollow)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Println("Feed created successfully!")
	fmt.Println("==========================================")
	printFeed(feed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("*ID				%s\n", feed.ID)
	fmt.Printf("*CreatedAt		%v\n", feed.CreatedAt)
	fmt.Printf("*UpdatedAt		%v\n", feed.UpdatedAt)
	fmt.Printf("*Name			%s\n", feed.Name)
	fmt.Printf("*Url			%s\n", feed.Url)
	fmt.Printf("*UserId			%v\n", feed.UserID)
}