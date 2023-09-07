package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tonrock01/quiz_go_pokeapi/dataModel"
)

var pokemonData dataModel.Pokemon
var berriesData dataModel.Pokemon

func GetDataPokeApi(apiURL string, endPoint string, dataStruct interface{}) {
	response, err := http.Get(apiURL + endPoint)
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
	if err := decoder.Decode(dataStruct); err != nil {
		fmt.Println("Can't convert to JSON:", err)
		return
	}
}

func main() {
	pokemonURL := "https://pokeapi.co/api/v2/pokemon/"
	berriesURL := "https://pokeapi.co/api/v2/berry/"

	GetDataPokeApi(pokemonURL, "30", &pokemonData)
	GetDataPokeApi(berriesURL, "20", &berriesData)

	fmt.Println("Pokemon Name:", pokemonData.Name)
	fmt.Println("Berry Name:", berriesData.Name)
}
