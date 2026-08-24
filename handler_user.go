package main

import (
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}

	fmt.Printf("User set to: %s\n", username)
	return nil
}