package main

import (
	"fmt"
	"log"
	"github.com/mumairdotdev/gator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	// Example usage of SetUser
	err = cfg.SetUser("Umair")
	if err != nil {
		log.Fatal("Error setting user:", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("Error reading config after setting user:", err)
	}
	fmt.Printf("Updated config: %+v\n", cfg)
}