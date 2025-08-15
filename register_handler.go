package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/frontendninja10/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	ctx := context.Background()

	username := cmd.args[0]

	_, err := s.db.GetUser(ctx, username)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if err == nil  {
		fmt.Fprintln(os.Stderr, "user already exists")
		os.Exit(1)
	}

	userArgs := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: username,
	}

	user, err := s.db.CreateUser(ctx, userArgs)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err = s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("User created successfully:", user.Name)
	return nil
}