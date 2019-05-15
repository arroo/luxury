package main

// Mover ...
type Mover interface {
	Move()
}

// GrabGems ...
type GrabGems struct {
	// player
	// gems
}

// Move ...
func (m *GrabGems) Move() {}

// Valid ...
func (m *GrabGems) Valid() bool { return true }

// ReserveCard ...
type ReserveCard struct {
	// player
	// card
}

// Move ...
func (m *ReserveCard) Move() {}

// Valid ...
func (m *ReserveCard) Valid() bool { return true }

// BuyCard ...
type BuyCard struct {
	// player
	// card
}

// Move ...
func (m *BuyCard) Move() {}

// Valid ...
func (m *BuyCard) Valid() bool { return true }
