package main

import (
	"errors"
	"fmt"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("a username is required")
	}

	username := cmd.args[0]

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("user has been set successfully: %s", username)
	return nil
}