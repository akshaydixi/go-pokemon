package main

import (
	"encoding/json"
	"fmt"
	"net/http"
//  "io/ioutil"
//  "strconv"
)

const (
	endpoint = "http://pokeapi.co/api/v1"
)

type Game struct {
  Name string `json:"name"`
  Id int `json:"id"`
  Resource_uri string `json:"resource_uri"`
  Created string `json:"created"`
  Modified string `json:"modified"`
  Release_year int `json:"release"year"`
  Generation int `json:"generation"`
}

type Ability struct {
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
}

type Description struct {
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
}

type EggGroup struct {
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
}

type Evolution struct {
	Level       int    `json:"level"`
	Method      string `json:"method"`
	Resouce_uri string `json:"resource_uri"`
	To          string `json:"to"`
}

type Move struct {
	Learn_type   string `json:"learn_type"`
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
	Level        int    `json:"level"`
}

type Sprite struct {
	Name        string `json:"name"`
	Resouce_uri string `json:"resource_uri"`
}

type Type struct {
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
}

type Pokemon struct {
	Name              string        `json:"name"`
	National_id       int           `json:"national_id"`
	Resource_uri      string        `json:"resource_uri"`
	Created           string        `json:"created"`
	Modified          string        `json:"modified"`
	Abilites          []Ability     `json:"abilities"`
	Egg_groups        []EggGroup    `json:"egg_groups"`
	Evolutions        []Evolution   `json:"evolutions"`
	Descriptions      []Description `json:"descriptions"`
	Moves             []Move        `json:"moves"`
	Types             []Type        `json:"types"`
	Catch_rate        int           `json:"catch_rate"`
	Species           string        `json:"species"`
	Hp                int           `json:"hp"`
	Attack            int           `json:"attack"`
	Defense           int           `json:"defense"`
	Sp_atk            int           `json:"sp_atk"`
	Sp_def            int           `json:"sp_def"`
	Speed             int           `json:"speed"`
	Egg_cycles        int           `json:"egg_cycles"`
	Ev_yield          string        `json:"ev_yield"`
	Exp               int           `json:"exp"`
	Growth_rate       string        `json:"growth_rate"`
	Happiness         int           `json:"happiness"`
	Height            string        `json:"height"`
	Male_female_ratio string        `json:"male_female_ratio"`
	Pkdx_id           int           `json:"pkdx_id"`
	Sprites           []Sprite      `json:"sprites"`
	Total             int           `json:"total"`
	Weight            string           `json:"weight"`
}

type Pokedex struct {
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
	Created      string `json:"created"`
	Modified     string `json:"modified"`
	Pokemon      []Pokemon `json:"pokemon"`
}

func main() {
  _, err := getPokemon("bulbasaur")
  if err != nil {
    fmt.Println(err)
  }
}


func getPokemon(identifier string) (Pokemon, error) {
  url := endpoint + "/pokemon/" + identifier
  res, err := http.Get(url)
  if err != nil {
    return Pokemon{}, err
  }
  defer res.Body.Close()
  decoder := json.NewDecoder(res.Body)
  var pokemon Pokemon
  if err := decoder.Decode(&pokemon); err != nil {
    return Pokemon{}, err
  }
  fmt.Println(pokemon.Name, pokemon.Species)
  return pokemon, nil
}

func getGame(identifier string) (Game, error) {
  url := endpoint + "/game/" + identifier
  res, err := http.Get(url)
  if err != nil {
    return Game{}, err
  }
  defer res.Body.Close()
  decoder := json.NewDecoder(res.Body)
  var game Game
  if err := decoder.Decode(&game); err != nil {
    return Game{}, err
  }
  fmt.Println(game.Name,":",game.Generation)
  return game, nil
}
