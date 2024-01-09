package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/andrejscek/pokedex/internal/pokeapi"
)

const (
	configFilename = "save.json"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

type savedConfig struct {
	NextLocationsURL *string                    `json:"nextLocationsURL"`
	PrevLocationsURL *string                    `json:"prevLocationsURL"`
	CaughtPokemon    map[string]pokeapi.Pokemon `json:"caughtPokemon"`
}

func saveConfigToFile(cfg *config, filename string) error {
	savedCfg := savedConfig{
		NextLocationsURL: cfg.nextLocationsURL,
		PrevLocationsURL: cfg.prevLocationsURL,
		CaughtPokemon:    cfg.caughtPokemon,
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(savedCfg); err != nil {
		return fmt.Errorf("error encoding and writing JSON data to file: %v", err)
	}

	return nil
}

func loadConfigFromFile(cfg *config, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var savedCfg savedConfig
	if err := decoder.Decode(&savedCfg); err != nil {
		return fmt.Errorf("error decoding JSON data: %v", err)
	}

	cfg.nextLocationsURL = savedCfg.NextLocationsURL
	cfg.prevLocationsURL = savedCfg.PrevLocationsURL
	cfg.caughtPokemon = savedCfg.CaughtPokemon

	return nil
}
