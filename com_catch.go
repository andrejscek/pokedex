package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) (bool, error) {

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return false, err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Trowing Pokebal to catch %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Println("Pokemon escaped!")
		return false, nil
	}

	fmt.Printf("%s caught! You can now inspect it with inspect command\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return true, nil

}
