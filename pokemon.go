package pokemon

import (
	"encoding/json"
	"net/http"
)

const (
	endpoint = "http://pokeapi.co/api/v1"
)

type Game struct {
	Name         string `json:"name"`
	Id           int    `json:"id"`
	Resource_uri string `json:"resource_uri"`
	Created      string `json:"created"`
	Modified     string `json:"modified"`
	Release_year int    `json:"release"year"`
	Generation   int    `json:"generation"`
}

type Ability struct {
	Name         string `json:"name"`
	Id           int    `json:"id"`
	Resource_uri string `json:"resource_uri"`
	Created      string `json:"created"`
	Modified     string `json:"modified"`
	Description  string `json:"description"`
}

type Description struct {
	Name         string  `json:"name"`
	Id           int     `json:"id"`
	Resource_uri string  `json:"resource_uri"`
	Created      string  `json:"created"`
	Modified     string  `json:"modified"`
	Games        []Game  `json:"games"`
	Pokemon      Pokemon `json:"pokemon"`
}

type EggGroup struct {
	Name         string    `json:"name"`
	Id           int       `json:"id"`
	Resource_uri string    `json:"resource_uri"`
	Created      string    `json:"created"`
	Modified     string    `json:"modified"`
	Pokemon      []Pokemon `json:"pokemon"`
}

type Evolution struct {
	Level       int    `json:"level"`
	Method      string `json:"method"`
	Resouce_uri string `json:"resource_uri"`
	To          string `json:"to"`
}

type Pokemon_Move struct {
	Learn_type   string `json:"learn_type"`
	Name         string `json:"name"`
	Resource_uri string `json:"resource_uri"`
	Level        int    `json:"level"`
}

type Move struct {
	Name         string `json:"name"`
	Id           int    `json:"id"`
	Resource_uri string `json:"resource_uri"`
	Created      string `json:"created"`
	Modified     string `json:"modified"`
	Description  string `json:"description"`
	Power        int    `json:"power"`
	Accuracy     int    `json:"accuracy"`
	Category     string `json:"category"`
	Pp           int    `json:"pp"`
}

type Sprite struct {
	Name        string  `json:"name"`
	Id          int     `json:"id"`
	Resouce_uri string  `json:"resource_uri"`
	Created     string  `json:"created"`
	Modified    string  `json:"modified"`
	Pokemon     Pokemon `json:"pokemon"`
	Image       string  `json:"image"`
}

type Type struct {
	Name            string `json:"name"`
	Id              int    `json:"id"`
	Resource_uri    string `json:"resource_uri"`
	Created         string `json:created"`
	Modified        string `json:"modified"`
	Ineffective     []Type `json:"ineffective"`
	No_effect       []Type `json:"no_effect"`
	Resistance      []Type `json:"resistance"`
	Super_effective []Type `json:"super_effective"`
	Weakness        []Type `json:"weakness"`
}

type Pokemon struct {
	Name              string         `json:"name"`
	National_id       int            `json:"national_id"`
	Resource_uri      string         `json:"resource_uri"`
	Created           string         `json:"created"`
	Modified          string         `json:"modified"`
	Abilites          []Ability      `json:"abilities"`
	Egg_groups        []EggGroup     `json:"egg_groups"`
	Evolutions        []Evolution    `json:"evolutions"`
	Descriptions      []Description  `json:"descriptions"`
	Moves             []Pokemon_Move `json:"moves"`
	Types             []Type         `json:"types"`
	Catch_rate        int            `json:"catch_rate"`
	Species           string         `json:"species"`
	Hp                int            `json:"hp"`
	Attack            int            `json:"attack"`
	Defense           int            `json:"defense"`
	Sp_atk            int            `json:"sp_atk"`
	Sp_def            int            `json:"sp_def"`
	Speed             int            `json:"speed"`
	Egg_cycles        int            `json:"egg_cycles"`
	Ev_yield          string         `json:"ev_yield"`
	Exp               int            `json:"exp"`
	Growth_rate       string         `json:"growth_rate"`
	Happiness         int            `json:"happiness"`
	Height            string         `json:"height"`
	Male_female_ratio string         `json:"male_female_ratio"`
	Pkdx_id           int            `json:"pkdx_id"`
	Sprites           []Sprite       `json:"sprites"`
	Total             int            `json:"total"`
	Weight            string         `json:"weight"`
}

type Pokedex struct {
	Name         string    `json:"name"`
	Resource_uri string    `json:"resource_uri"`
	Created      string    `json:"created"`
	Modified     string    `json:"modified"`
	Pokemon      []Pokemon `json:"pokemon"`
}

// This function gets the JSON from the API and populates the value field
// which is passed by reference to it
func endpointRequest(url string, value interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(&value)
}

func GetPokedex(identifier string) (pokedex Pokedex, err error) {
	url := endpoint + "/pokedex/" + identifier
	if err = endpointRequest(url, &pokedex); err != nil {
		return Pokedex{}, err
	}
	return pokedex, nil
}
func GetPokemon(identifier string) (pokemon Pokemon, err error) {
	url := endpoint + "/pokemon/" + identifier
	if err = endpointRequest(url, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}

func GetGame(identifier string) (game Game, err error) {
	url := endpoint + "/game/" + identifier
	if err = endpointRequest(url, &game); err != nil {
		return Game{}, err
	}
	return game, nil
}

func GetType(identifier string) (type_ Type, err error) {
	url := endpoint + "/type/" + identifier
	if err = endpointRequest(url, &type_); err != nil {
		return Type{}, err
	}
	return type_, nil
}

func GetMove(identifier string) (move Move, err error) {
	url := endpoint + "/move/" + identifier
	if err = endpointRequest(url, &move); err != nil {
		return Move{}, err
	}
	return move, nil
}

func GetAbility(identifier string) (ability Ability, err error) {
	url := endpoint + "/ability/" + identifier
	if err = endpointRequest(url, &ability); err != nil {
		return Ability{}, err
	}
	return ability, nil
}

func GetEggGroup(identifier string) (eggGroup EggGroup, err error) {
	url := endpoint + "/egg/" + identifier
	if err = endpointRequest(url, &eggGroup); err != nil {
		return EggGroup{}, err
	}
	return eggGroup, nil
}

func GetDescription(identifier string) (description Description, err error) {
	url := endpoint + "/description/" + identifier
	if err = endpointRequest(url, &description); err != nil {
		return Description{}, err
	}
	return description, nil
}

func GetSprite(identifier string) (sprite Sprite, err error) {
	url := endpoint + "/sprite/" + identifier
	if err = endpointRequest(url, &sprite); err != nil {
		return Sprite{}, err
	}
	return sprite, nil
}
