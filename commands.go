package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(s *state, cmd command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdHandler, exists := c.registeredCommands[cmd.name]
	if !exists {
		return errors.New("command not found")
	}
	return cmdHandler(s, cmd)
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.registeredCommands[name] = f
}