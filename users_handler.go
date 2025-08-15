package main

import (
	"context"
	"fmt"
)

func getUsers(s *state, cmd command) error {
	ctx := context.Background()

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch users: %w", err)
	}

	currentUser := s.cfg.CurrentUsername

	for _, user := range users {
		if currentUser == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Println("*", user.Name)
	}
	return nil
}