package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
  "strconv"
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
	Species           string        `json:"specied"`
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
	Weight            int           `json:"weight"`
}

type Pokedex struct {
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
	Created      string `json:"created"`
	Modified     string `json:"modified"`
	Pokemon      []Pokemon `json:"pokemon"`
}

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func get_(resource_id string, val_1 string, val_2 int ) interface{} {
  url := endpoint + "/" + resource_id + "/"
  if (val_1 == "NULL") {
    url = url + strconv.Itoa(val_2)
  } else {
    url = url + val_1
  }
  fmt.Println(url)
  res, err := http.Get(url)
  perror(err)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  switch resource_id {
    case "pokemon": {var resource Pokemon; json.Unmarshal(body, &resource); return resource;}
    case "move": {var resource Move; json.Unmarshal(body, &resource); return resource;}
    case "pokedex": {var resource Pokedex; json.Unmarshal(body, &resource); return resource;}
    case "type" : {var resource Type; json.Unmarshal(body, &resource); return resource;}
    case "ability" : {var resource Ability; json.Unmarshal(body,&resource); return resource;}
    case "egg" : {var resource EggGroup; json.Unmarshal(body, &resource); return resource;}
    case "description" : {var resource Description; json.Unmarshal(body,&resource); return resource;}
    case "sprite" : {var resource Sprite; json.Unmarshal(body, &resource); return resource;}
    case "game" : {var resource Game; json.Unmarshal(body, &resource); return resource;}
  }
  //json.Unmarshal(body, &resource)
  //fmt.Println("%v\n",resource)
return 
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
	get_1("pokemon", 2)
}

func get(resource_id string, value int) interface{} {
  return get_(resource_id, "NULL", value)
}
