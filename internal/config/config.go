package config

import (
	"strings"

	"github.com/gaga4l/pokedex/internal/pokeapi"
)

type Config struct {
	Commands      map[string]CliCommand
	Pokedex       map[string]pokeapi.PokemonResponse
	PokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	BaseURL       string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

func CleanInput(text string) []string {
	var result []string
	text = strings.ToLower(text)
	result = strings.Fields(text)

	return result
}
