package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mumairdotdev/gator-go/internal/config"
	"github.com/mumairdotdev/gator-go/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	dbQueries := database.New(db)

	State := state{cfg: &cfg, db: dbQueries}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleGetUsers)

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
