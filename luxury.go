package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	tier1, err := getDeck("tier1.json")
	if err != nil {
		fmt.Printf("error decking: %v", err)

		return
	}

	fmt.Printf("cards: %v\n", tier1)

	nobles, err := getNobles("nobles.json")
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
