package main

import (
	"fmt"
	"math/rand"
)

// Game is the game of Luxury
type Game struct {
	Nobles Nobles
	Tier1  Deck
	Tier2  Deck
	Tier3  Deck
	Gems   map[Gem]int
}

// Gems2Player ...
var Gems2Player = map[Gem]int{
	Gold:     5,
	Diamond:  4,
	Emerald:  4,
	Jet:      4,
	Ruby:     4,
	Sapphire: 4,
}

// Gems3Player ...
var Gems3Player = map[Gem]int{
	Gold:     5,
	Diamond:  5,
	Emerald:  5,
	Jet:      5,
	Ruby:     5,
	Sapphire: 5,
}

// Gems4Player ...
var Gems4Player = map[Gem]int{
	Gold:     5,
	Diamond:  7,
	Emerald:  7,
	Jet:      7,
	Ruby:     7,
	Sapphire: 7,
}

// NewGame inits a new game of Luxury
func NewGame(nobles Nobles, t1, t2, t3 Deck, gems map[Gem]int) *Game {
	return &Game{
		Nobles: nobles,
		Tier1:  t1,
		Tier2:  t2,
		Tier3:  t3,
		Gems:   gems,
	}
}

func (g *Game) shuffle() {
	rand.Shuffle(len(g.Nobles), g.Nobles.Swapper())
	rand.Shuffle(len(g.Tier1), g.Tier1.Swapper())
	rand.Shuffle(len(g.Tier2), g.Tier2.Swapper())
	rand.Shuffle(len(g.Tier3), g.Tier3.Swapper())
}

// Start starts a game of luxury
func (g *Game) Start(players []*Player) {
	g.shuffle()

	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

turn:
	for {
		for _, p := range players {
			p.Move(g)
		}

		// check if last turn has happened
		for _, p := range players {
			// game is over once everyone has had a turn and someone is at at least 15 points
			if p.VictoryPoints() >= 15 {
				break turn
			}
		}
	}

	wI := 0
	for i, p := range players {
		if p.VictoryPoints() > players[wI].VictoryPoints() {
			wI = i
		}
	}

	fmt.Printf("Winner: %s\n", players[wI].Name)
}

// BoardState ...
type BoardState struct {
	Players      []Player
	Nobles       []Noble
	VisibleCards [][]Card
}

// StartWithChans ...
func (g *Game) StartWithChans(p []chan Mover) {
	g.shuffle()

	rand.Shuffle(len(p), func(i, j int) {
		p[i], p[j] = p[j], p[i]
	})
}
