package main

// Card is a game piece
type Card struct {
	Provides      Gem         `json:"gem"`
	VictoryPoints int         `json:"vp"`
	Cost          map[Gem]int `json:"cost"`
}

// Deck is an ordered collection of cards
type Deck []Card

// Swapper returns a swap function suitable for rand.Shuffle
func (d Deck) Swapper() func(int, int) {
	return func(i, j int) {
		d[i], d[j] = d[j], d[i]
	}
}
