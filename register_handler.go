package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/frontendninja10/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("you must pass in a name to register")
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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