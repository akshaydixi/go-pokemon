package pokemon

import (
	"testing"
)

func Test_GamePokemonRed(t *testing.T) {
	game, err := getGame("4")
	if err != nil {
		t.Error("Error: ", err)
	} else {
		t.Log("Passed: ", game.Name)
	}
}

func Test_GameError(t *testing.T) {
	game, err := getGame("1231")
	if err != nil {
		t.Log("Passed: ", err)
	} else {
		t.Error("Error: ", game)
	}
}

func Test_PokemonCharizardByName(t *testing.T) {
	pokemon, err := getPokemon("charizard")
	if err != nil {
		t.Error("Error: ", err)
	} else if pokemon.Name != "Charizard" {
		t.Error("Error: Wrong Pokemon???")
	} else {
		t.Log("Passed:", pokemon.Name, pokemon.National_id)
	}
}

func Test_PokemonCharizardById(t *testing.T) {
	pokemon, err := getPokemon("6")
	if err != nil {
		t.Error("Error: ", err)
	} else if pokemon.Name != "Charizard" {
		t.Error("Error: Wrong Pokemon???")
	} else {
		t.Log("Passed: ", pokemon.Name, pokemon.National_id)
	}
}

func Test_PokedexCallSuccess(t *testing.T) {
	pokedex, err := getPokedex("1")
	if err != nil {
		t.Error("Error: ", err)
	} else {
		t.Log("Passed :", pokedex.Name, " contains ", len(pokedex.Pokemon), " pokemon")
	}
}
