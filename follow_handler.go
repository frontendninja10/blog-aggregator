package main

import (
	"context"
	"fmt"
	"time"

	"github.com/frontendninja10/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func followHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.name)
	}

	ctx := context.Background()

	feed, err := s.db.GetFeed(ctx, cmd.args[0])
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

	_, err = s.db.CreateFeedFollow(ctx, createFeedFollow)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	return nil
}

func unfollowHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.name)
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

	deleteFeedFollowArg := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.db.DeleteFeedFollow(context.Background(), deleteFeedFollowArg); err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Printf("%s unfollowed successfully\n", feed.Name)

	return nil
}


func listFeedFollowsHandler(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
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