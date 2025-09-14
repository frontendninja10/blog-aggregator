package users

import (
	"context"
	"fmt"

	"github.com/frontendninja10/blog-aggregator/internal/app"
)

func Reset(s *app.State, cmd app.Command) error {
	ctx := context.Background()

	if err := s.DB.DeleteUsers(ctx); err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	s.Config.SetUser("")
	fmt.Println("Database reset successfully!")
	return nil
}