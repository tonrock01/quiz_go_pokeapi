package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
}

type Berries struct {
	Name string `json:"name"`
}

var pokemonData Pokemon
var berriesData Pokemon

func GetDataPokeApi(apiURL string, data interface{}) {
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Bad request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println(response.Status)
		return
	}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(data); err != nil {
		fmt.Println("Can't convert to JSON:", err)
		return
	}
}

func main() {
	pokemonURL := "https://pokeapi.co/api/v2/pokemon/"
	berriesURL := "https://pokeapi.co/api/v2/berry/"

	GetDataPokeApi(pokemonURL+"1", &pokemonData)
	GetDataPokeApi(berriesURL+"10", &berriesData)

	fmt.Println("ชื่อ Pokemon:", pokemonData.Name)
	fmt.Println("ชื่อ Berry:", berriesData.Name)
}
