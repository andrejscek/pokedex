package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Trowing Pokebal to catch %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Println("Pokemon escaped!")
		return nil
	}

	fmt.Printf("%s caught! You can now inspect it with inspect command\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil

}
