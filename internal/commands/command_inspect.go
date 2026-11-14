package commands

import (
	"errors"
	"fmt"

	"github.com/gaga4l/pokedex/internal/config"
)

func CommandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please enter the name of pokemon")
	}

	pokemonName := args[0]

	pokemon, exist := cfg.Pokedex[pokemonName]
	if !exist {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Printf("Stats: \n")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}
