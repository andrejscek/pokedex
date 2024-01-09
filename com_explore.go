package main

func commandExplore(cfg *config, name string) ([]string, error) {

	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, enc := range location.PokemonEncounters {
		names = append(names, enc.Pokemon.Name)
	}
	return names, nil
}
