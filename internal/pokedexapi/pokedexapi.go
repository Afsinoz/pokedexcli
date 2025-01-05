package pokedexapi

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

func MapGet(URL string) (next string, previous string, results []string, err error) {
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
