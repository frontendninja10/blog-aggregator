package users

import (
	"context"
	"fmt"

	"github.com/frontendninja10/blog-aggregator/internal/app"
)

func ListUsers(s *app.State, cmd app.Command) error {
	ctx := context.Background()

	users, err := s.DB.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch users: %w", err)
	}

	currentUser := s.Config.CurrentUsername

	for _, user := range users {
		if currentUser == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Println("*", user.Name)
	}
	return nil
}