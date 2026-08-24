package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if cmd.Name == "login" {
		if handler, ok := c.registeredCommands[cmd.Name]; ok {
			return handler(s, cmd)
		}
	}
	return errors.New("command not found")
}

func (c *commands) register(name string, f func(*state, command) error) {
	if name == "login" {
		c.registeredCommands[name] = f
	}
}
