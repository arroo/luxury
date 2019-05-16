package main

// Mover ...
type Mover interface {
	Move(*Game) *Game
}

// GrabGems ...
type GrabGems struct {
	// player
	// gems
}

// Move ...
func (m *GrabGems) Move(g *Game) *Game {
	return g
}

// Valid ...
func (m *GrabGems) Valid() bool { return true }

// ReserveCard ...
type ReserveCard struct {
	// player
	// card
}

// Move ...
func (m *ReserveCard) Move(g *Game) *Game {
	return g
}

// Valid ...
func (m *ReserveCard) Valid() bool { return true }

// BuyCard ...
type BuyCard struct {
	// player
	// card
}

// Move ...
func (m *BuyCard) Move(g *Game) *Game {
	return g
}

// Valid ...
func (m *BuyCard) Valid() bool { return true }
