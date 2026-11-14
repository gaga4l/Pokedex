package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/gaga4l/pokedex/internal/config"
)

func CommandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please enter the name of pokemon you want to catch")
	}

	pokemonName := args[0]
	pokemon, err := cfg.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catchChance := 100 / (1 + pokemon.BaseExperience/100)

	if rand.Intn(100) >= catchChance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.Pokedex[pokemon.Name] = pokemon

	return nil
}
