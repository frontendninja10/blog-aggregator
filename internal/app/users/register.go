package users

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/frontendninja10/blog-aggregator/internal/app"
	"github.com/frontendninja10/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func Register(s *app.State, cmd app.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	ctx := context.Background()

	username := cmd.Args[0]

	_, err := s.DB.GetUser(ctx, username)
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

	user, err := s.DB.CreateUser(ctx, userArgs)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err = s.Config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("User created successfully:", user.Name)
	return nil
}