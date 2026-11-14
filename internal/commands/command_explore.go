package commands

import (
	"fmt"

	"github.com/gaga4l/pokedex/internal/config"
)

func CommandExplore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("enter location-area name")
	}
	locationName := args[0]
	resp, err := cfg.PokeapiClient.GetFromLocation(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
