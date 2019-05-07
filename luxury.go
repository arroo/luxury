package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	for i := 1; i < 4; i++ {
		tier, err := getDeck(fmt.Sprintf("cards/tier%d.json", i))
		if err != nil {
			fmt.Printf("error decking: %v", err)

			return
		}

		fmt.Printf("\ncards tier%d: %v\n", i, tier)
	}

	nobles, err := getNobles("nobles/nobles.json")
	if err != nil {
		fmt.Printf("error nobling: %v", err)

		return
	}

	fmt.Printf("nobles: %v\n", nobles)
}

func getNobles(fileName string) (Nobles, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	nobleCollection := struct {
		Nobles Nobles `json:"nobles"`
	}{}

	err = json.Unmarshal([]byte(file), &nobleCollection)
	if err != nil {
		return nil, err
	}

	return nobleCollection.Nobles, nil
}

func getDeck(fileName string) (Deck, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	deck := struct {
		Deck Deck `json:"cards"`
	}{}

	err = json.Unmarshal([]byte(file), &deck)
	if err != nil {
		return nil, err
	}

	return deck.Deck, nil
}
