package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("https://pokeapi.co/api/v2/location/?limit=20&offset=0")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

}
