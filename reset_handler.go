package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	ctx := context.Background()

	if err := s.db.DeleteUsers(ctx); err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	s.cfg.SetUser("")
	fmt.Println("Database reset successfully!")
	return nil
}