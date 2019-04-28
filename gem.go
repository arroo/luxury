package main

import (
	"encoding/json"
	"fmt"
)

// Gem is the currency of this game
type Gem int

// gems
const (
	Gold Gem = iota
	Diamond
	Emerald
	Jet
	Ruby
	Sapphire
)

const (
	goldStr     = "gold"
	diamondStr  = "diamond"
	emeraldStr  = "emerald"
	jetStr      = "jet"
	rubyStr     = "ruby"
	sapphireStr = "sapphire"
)

// Colour representation of each type of gem
type Colour int

// Colours of each gem
const (
	White Colour = iota
	Green
	Black
	Red
	Blue
)

// Token is a valued game piece
type Token struct {
	Kind Gem
}

var gemToString = map[Gem]string{
	Gold:     goldStr,
	Diamond:  diamondStr,
	Emerald:  emeraldStr,
	Jet:      jetStr,
	Ruby:     rubyStr,
	Sapphire: sapphireStr,
}

var stringToGem = map[string]Gem{
	goldStr:     Gold,
	diamondStr:  Diamond,
	emeraldStr:  Emerald,
	jetStr:      Jet,
	rubyStr:     Ruby,
	sapphireStr: Sapphire,
}

var gemToColour = map[Gem]Colour{
	Diamond:  White,
	Emerald:  Green,
	Jet:      Black,
	Ruby:     Red,
	Sapphire: Blue,
}

func (g Gem) String() string {
	s, ok := gemToString[g]
	if !ok {
		return "Unknown"
	}

	return s
}

// MarshalJSON marshals the enum as a quoted json string
func (g Gem) MarshalJSON() ([]byte, error) {
	s, ok := gemToString[g]
	if !ok {
		return nil, fmt.Errorf("unmapped gem: %v", g)
	}

	return json.Marshal(s)
}

// UnmarshalJSON unmarshals a quoted json string to the enum value
func (g *Gem) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	var ok bool
	*g, ok = stringToGem[j]
	if !ok {
		return fmt.Errorf("unmapped gem: %s", j)
	}

	return nil
}

// MarshalText marshals the enum as a string
func (g Gem) MarshalText() ([]byte, error) {
	s, ok := gemToString[g]
	if !ok {
		return nil, fmt.Errorf("unmapped gem: %v", g)
	}

	return []byte(s), nil
}

// UnmarshalText unmarshals a string to the enum value
func (g *Gem) UnmarshalText(b []byte) error {
	var j = string(b)

	var ok bool
	*g, ok = stringToGem[j]
	if !ok {
		return fmt.Errorf("unmapped gem: %s", j)
	}

	return nil
}
