package main

import (
	"time"

	"github.com/gaga4l/pokedex/internal/config"
	"github.com/gaga4l/pokedex/internal/pokeapi"
)

const baseURL = "https://pokeapi.co/api/v2/"

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	commands := getCommands()
	pokedex := map[string]pokeapi.PokemonResponse{}
	cfg := &config.Config{
		PokeapiClient: *pokeClient,
		Commands:      commands,
		Pokedex:       pokedex,
		BaseURL:       baseURL,
	}

	startRepl(cfg)
}
