package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("a username is required")
	}
	
	ctx := context.Background()

	username := cmd.args[0]

	_, err := s.db.GetUser(ctx, username)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if err == sql.ErrNoRows {
		fmt.Fprintln(os.Stderr, "user does not exist")
		os.Exit(1)
	}

	if s.cfg.CurrentUsername == username {
		fmt.Fprintln(os.Stderr, "user already logged in")
		os.Exit(1)
	}

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Println("user has been set successfully:", username)
	return nil
}