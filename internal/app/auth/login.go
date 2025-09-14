package auth

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/frontendninja10/blog-aggregator/internal/app"
)

func Login(s *app.State, cmd app.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	
	ctx := context.Background()

	username := cmd.Args[0]

	_, err := s.DB.GetUser(ctx, username)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if err == sql.ErrNoRows {
		fmt.Fprintln(os.Stderr, "user does not exist")
		os.Exit(1)
	}

	if s.Config.CurrentUsername == username {
		fmt.Fprintln(os.Stderr, "user already logged in")
		os.Exit(1)
	}

	if err := s.Config.SetUser(username); err != nil {
		return err
	}

	fmt.Println("user has been set successfully:", username)
	return nil
}