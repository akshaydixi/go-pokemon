package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"

)

const (
  endpoint = "http://pokeapi.co/api/v1"
)
type Ability struct {
  Name string `json:"name"`
  Uri string  `json:"resource_uri"`
}

type Description struct {
  Name string `json:"name"`
  Uri string `json:"resouce_uri"`
}

type EggGroup struct {
  Name string`json:"name"`
  ResourceUri string `json:"resource_uri"`
}

type Evolution struct {
  Level int `json:"level"`
  Method string `json:"method"`
  Uri string `json:"resource_uri"`
  To string `json:"to"`
}

type Move struct {
  Learn_type string `json:"learn_type"`
  Name string`json:"name"`
  Uri string `json:"resource_uri"`
  Level int `json:"level"`
}

type Sprite struct {
  Name string`json:"name"`
  Uri string `json:"resource_uri"`
}

type Type struct {
  Name string`json:"name"`
  Uri string `json:"resource_uri"`
}

type Pokemon struct {
  Name string`json:"name"`
  National_id int `json:"national_id"`
  Uri string `json:"resource_uri"`
  Created string `json:"created"`
  Modified string `json:"modified"`
  Abilites []Ability `json:"abilities"`
  Egg_groups []EggGroup `json:"egg_groups"`
  Evolutions []Evolution `json:"evolutions"`
  Descriptions []Description `json:"descriptions"`
  Moves []Move `json:"moves"`
  Types []Type `json:"types"`
  Catch_rate int `json:"catch_rate"`
  Species string `json:"specied"`
  Hp int `json:"hp"`
  Attack int `json:"attack"`
  Defense int `json:"defense"`
  Sp_atk int `json:"sp_atk"`
  Sp_def int `json:"sp_def"`
  Speed int `json:"speed"`
  Egg_cycles int `json:"egg_cycles"`
  Ev_yield string `json:"ev_yield"`
  Exp int `json:"exp"`
  Growth_rate string `json:"growth_rate"`
  Happiness int `json:"happiness"`
  Height string `json:"height"`
  Male_female_ratio string `json:"male_female_ratio"`
  Pkdx_id int `json:"pkdx_id"`
  Sprites []Sprite  `json:"sprites"`
  Total int `json:"total"`
  Weight int  `json:"weight"`
}

type Pokedex struct {
  Name string `json:"name"`
  Resource_uri string `json:"resource_uri"`
  Created string  `json:"created"`
  Modified string `json:"modified"`
  Pokemon []Pokemon `json:"pokemon"`
}
func perror(err error) {
  if err != nil {
    panic(err)
  }
}

func get_content() {
  url := endpoint + "/pokemon/1"
  res, err := http.Get(url)
  perror(err)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  perror(err)
  var pokemon Pokemon
  json.Unmarshal(body, &pokemon)
  fmt.Printf("Results: %v\n", pokemon)
}

func main() {
  get_content()
}
