package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gaga4l/pokedex/internal/commands"
	"github.com/gaga4l/pokedex/internal/config"
)

func startRepl(cfg *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		commands := config.CleanInput(scanner.Text())
		if len(commands) == 0 {
			continue
		}

		commandName := commands[0]
		args := []string{}
		if len(commands) > 1 {
			args = commands[1:]
		}

		cmd, ok := cfg.Commands[commandName]
		if ok {
			err := cmd.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
			continue

		}

	}
}

func getCommands() map[string]config.CliCommand {
	return map[string]config.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commands.CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    commands.CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "show 20 pokemon location areas",
			Callback:    commands.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "show previous 20 pokemon location areas",
			Callback:    commands.CommandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "explore pokemons in particular location areas",
			Callback:    commands.CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "try to catch a pokemon",
			Callback:    commands.CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "pokemon list in pokedex",
			Callback:    commands.CommandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "show infos of particular pokemon",
			Callback:    commands.CommandInspect,
		},
	}
}
