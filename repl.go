package main

import (
	"fmt"
	"os"

	"github.com/paulrademacher/climenu"
)

func exploreLoc(cfg *config, location string) {
	fmt.Println()
	pokes, err := commandExplore(cfg, location)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		menu := climenu.NewButtonMenu("Explore", "Select pokemon to try and catch it, Esc to go back")
		for _, k := range pokes {
			menu.AddMenuItem(k, k)
		}
		pokemon, escaped := menu.Run()
		if escaped {
			return
		}
		commandCatch(cfg, pokemon)
		fmt.Println()
	}

}

func exploreMap(cfg *config) {
	locations, err := commandMapf(cfg)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		return
	}
	for {
		menu := climenu.NewButtonMenu("Explore", "Select location to explore it, Esc to go back")
		menu.AddMenuItem("Next page", "map")
		menu.AddMenuItem("Previous page", "mapb")
		for _, k := range locations {
			menu.AddMenuItem(k, k)
		}
		selection, escaped := menu.Run()
		if escaped {
			break
		}
		switch selection {
		case "map":
			new_loc, err := commandMapf(cfg)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			locations = new_loc
		case "mapb":
			new_loc, err := commandMapb(cfg)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			locations = new_loc
		default:
			exploreLoc(cfg, selection)
		}

	}
}

func startRepl(cfg *config) {

	for {

		menu := climenu.NewButtonMenu("\n\nPokedex main menu", "Enter to choose an action, Esc to exit")
		menu.AddMenuItem("My Pokemons", "pokedex")
		menu.AddMenuItem("Explore", "map")
		menu.AddMenuItem("Explore custom location", "explore")
		menu.AddMenuItem("Save", "save")
		menu.AddMenuItem("Load", "load")

		action, escaped := menu.Run()
		if escaped {
			os.Exit(0)
		}

		switch action {
		case "pokedex":

			if len(cfg.caughtPokemon) == 0 {
				fmt.Println()
				fmt.Println("You have not caught any pokemon yet")
				fmt.Println()
				continue
			}

			menu := climenu.NewButtonMenu("My Pokemons", "Select pokemon to inspect it, Esc to go back")
			for _, k := range cfg.caughtPokemon {
				menu.AddMenuItem(k.Name, k.Name)
			}
			pokemon, escaped := menu.Run()
			if escaped {
				continue
			}
			commandInspect(cfg, pokemon)
			fmt.Println()

		case "map":
			exploreMap(cfg)

		case "explore":
			response := climenu.GetText("Enter custom location to explore", "location name")
			exploreLoc(cfg, response)

		case "save":
			err := saveConfigToFile(cfg, configFilename)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Saved!")
			}
		case "load":
			err := loadConfigFromFile(cfg, configFilename)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Loaded!")
			}
		}
	}
}
