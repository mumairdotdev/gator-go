package main

import (
	"log"
	"os"
	"github.com/mumairdotdev/gator-go/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {


	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}
	State := state{cfg: &cfg}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handleLogin)

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("No command provided")
	}

	cmd := command{
		Name: args[0],
		Args: args[1:],
	}

	err = cmds.run(&State, cmd)
	if err != nil {
		log.Fatal("Error executing command:", err)
	}
}