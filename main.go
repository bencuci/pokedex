package main

import (
	"github.com/bencuci/pokedex/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
}

func main() {
	pokeClient := pokeapi.NewClient(10*time.Second, 10*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
