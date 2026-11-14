package commands

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gaga4l/pokedex/internal/config"
	"github.com/gaga4l/pokedex/internal/pokeapi"
	"github.com/gaga4l/pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Minute)

func CommandMap(cfg *config.Config, args ...string) error {
	fullURL := cfg.BaseURL + "location-area/"
	if cfg.Next == nil {
		cfg.Next = &fullURL
	}
	pokeResponse, err := getPokeResponse(cfg, cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = pokeResponse.Next
	cfg.Previous = pokeResponse.Previous

	for _, result := range pokeResponse.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func CommandMapB(cfg *config.Config, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	pokeResponse, err := getPokeResponse(cfg, cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = cfg.Previous
	cfg.Previous = pokeResponse.Previous

	for _, result := range pokeResponse.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func getPokeResponse(cfg *config.Config, addr *string) (pokeapi.PokeApiResponse, error) {
	var pokeResponse pokeapi.PokeApiResponse
	cacheData, exist := cache.Get(*addr)
	if exist {
		if err := json.Unmarshal(cacheData, &pokeResponse); err != nil {
			return pokeapi.PokeApiResponse{}, err
		}
	} else {
		locations, err := cfg.PokeapiClient.ListLocations(addr)
		if err != nil {
			return pokeapi.PokeApiResponse{}, err
		}

		cacheData, err := json.Marshal(locations)
		if err != nil {
			return pokeapi.PokeApiResponse{}, err
		}

		cache.Add(*addr, cacheData)

		pokeResponse = locations
	}
	return pokeResponse, nil
}
