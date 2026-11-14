package commands

import (
	"errors"
	"fmt"

	"github.com/gaga4l/pokedex/internal/config"
)

func CommandPokedex(cfg *config.Config, args ...string) error {
	if len(cfg.Pokedex) == 0 {
		return errors.New("Your Pokedex is empty. \ntry to catch new pokemon with 'catch' command")
	}

	fmt.Printf("Your Pokedex: \n")

	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("  -%s\n", pokemon.Name)
	}

	return nil
}
