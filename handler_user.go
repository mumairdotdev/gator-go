package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mumairdotdev/gator-go/internal/database"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil || username != user.Name {
		return fmt.Errorf("user not found")
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}

	fmt.Printf("User set to: %s\n", username)
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	// Create a new user in the database
	newUser, err := s.db.CreatUser(context.Background(), database.CreatUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	err = s.cfg.SetUser(newUser.Name)
	if err != nil {
		return fmt.Errorf("failed to set user in config: %v", err)
	}
	fmt.Printf("User registered: %s (ID: %s)\n", newUser.Name, newUser.ID)
	return nil
}

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete all users: %v", err)
	}
	fmt.Println("All users deleted successfully.")
	return nil
}

func handleGetUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}
	for _, user := range users {
		// Ensure that the (current) follows the currently logged-in user.
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}