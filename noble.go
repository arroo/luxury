package main

// Noble is a type of bonus-granting game piece
type Noble struct {
	VictoryPoints int         `json:"vp"`
	Gems          map[Gem]int `json:"gems"`
}

// Nobles ordered collection of Nobles
type Nobles []Noble

// Swapper returns a swap function suitable for rand.Shuffle
func (n Nobles) Swapper() func(int, int) {
	return func(i, j int) {
		n[i], n[j] = n[j], n[i]
	}
}
