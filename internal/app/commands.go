package app

import (
	"errors"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	RegisteredCommands map[string]func(s *State, cmd Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	cmdHandler, exists := c.RegisteredCommands[cmd.Name]
	if !exists {
		return errors.New("command not found")
	}
	return cmdHandler(s, cmd)
}

func (c *Commands) Register(name string, f func(s *State, cmd Command) error) {
	c.RegisteredCommands[name] = f
}

func NewCommands() Commands {
	return Commands{
		RegisteredCommands: make(map[string]func(s *State, cmd Command) error),
	}
}