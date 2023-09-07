package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tonrock01/quiz_go_pokeapi/dataModel"
)

var pokemonData dataModel.Pokemon
var berriesData dataModel.Berries
var evolutionData dataModel.Evolution
var gamesData dataModel.Games

func GetDataPokeApi(apiURL string, endPoint string, dataStruct interface{}) {
	response, err := http.Get(apiURL + endPoint)
	if err != nil {
		fmt.Println("Bad request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Status Code:", response.Status)
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
	evolutionURL := "https://pokeapi.co/api/v2/evolution-chain/"
	gamesURL := "https://pokeapi.co/api/v2/generation/"

	GetDataPokeApi(pokemonURL, "6", &pokemonData)
	GetDataPokeApi(berriesURL, "12", &berriesData)
	GetDataPokeApi(evolutionURL, "10", &evolutionData)
	GetDataPokeApi(gamesURL, "1", &gamesData)

	fmt.Printf("Pokemon Name : %s\n -Species : %s\n", pokemonData.Name, pokemonData.Types[0].Type.Name)
	for i := 0; i < len(pokemonData.Abilities); i++ {
		fmt.Printf(" -Ability %d : %s\n", i+1, pokemonData.Abilities[i].Ability.Name)
	}

	fmt.Printf("Berry Name : %s\n -Flavors : %s\n", berriesData.Name, berriesData.Flavors[0].Flavor.Name)
	fmt.Printf("Pokemon Name : %s ---> Evole to : %s\n", evolutionData.Chain.Species.Name, evolutionData.Chain.EvolvesTo[0].Species.Name)
	fmt.Printf("Pokemon Game series : %s\n -Main Region : %s\n", gamesData.Name, gamesData.MainRegion.Name)
}
