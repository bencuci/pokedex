package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, argument string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore [location_name]",
			description: "Lists all of the pokemons at the given location",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config, argument string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, argument string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: ", cmd.name)
		fmt.Println(cmd.description)
	}

	return nil
}

func commandMapf(cfg *config, argument string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationsResp.Next
	cfg.prevURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, argument string) error {
	if cfg.prevURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	cfg.nextURL = cfg.prevURL

	return commandMapf(cfg, argument)
}

func commandExplore(cfg *config, locationName string) error {
	fmt.Println("Exploring " + locationName + "...")
	encountersResp, err := cfg.pokeapiClient.ListEncounters(locationName)
	if err != nil {
		return err
	}

	for _, encounter := range encountersResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
