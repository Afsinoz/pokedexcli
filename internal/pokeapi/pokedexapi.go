package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Afsinoz/pokedexcli/internal/pokecache"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Area struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func MapGet(URL string) (next string, previous string, results []string, err error) {
	//TODO: After finishing it make the cache 5 seconds, or smt less than 10
	cache := pokecache.NewCache(100 * time.Second)
	var data []byte
	byteValue, ok := cache.Get(URL)
	if ok {
		data = byteValue
	} else {

		resp, err := http.Get(URL)

		if err != nil {
			return "", "", []string{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", "", []string{}, err
		}

		cache.Add(URL, data)

	}

	var location Location

	if err := json.Unmarshal(data, &location); err != nil {
		return "", "", []string{}, err
	}

	places := []string{}

	for _, result := range location.Results {

		places = append(places, result.Name)

	}

	//	fmt.Println("The next url is : ", location.Next, "this is gonna be pointer")
	//
	//	fmt.Println("The previous url is : ", location.Previous, "this is gonna be pointer as well")
	//
	//	fmt.Println("The current url is : ", URL)
	if location.Previous == nil {
		return *location.Next, "", places, nil
	}

	return *location.Next, *location.Previous, places, nil

}

func PokeGet(URL string, nameArea string) ([]string, error) {

	//TODO: Change the caching time at the end
	cache := pokecache.NewCache(100 * time.Second)
	var data []byte

	fullURL := URL + nameArea

	byteValue, ok := cache.Get(fullURL)
	if ok {
		data = byteValue
	} else {

		resp, err := http.Get(fullURL)

		if err != nil {
			return []string{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return []string{}, err
		}
		cache.Add(fullURL, data)
	}
	var area Area

	if err := json.Unmarshal(data, &area); err != nil {
		return []string{}, err
	}

	pokemonNames := []string{}

	for _, pokemon := range area.PokemonEncounters {
		pokemonNames = append(pokemonNames, pokemon.Pokemon.Name)
	}
	return pokemonNames, nil

}
