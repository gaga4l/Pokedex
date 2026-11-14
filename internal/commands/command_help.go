package commands

import (
	"fmt"

	"github.com/gaga4l/pokedex/internal/config"
)

func CommandHelp(cfg *config.Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for cmd, val := range cfg.Commands {
		fmt.Printf("%s: %s\n", cmd, val.Description)
	}

	return nil
}
