package main

import (
	"fmt"
	"math/rand"
)

// Player is the entity playing Luxury
type Player struct {
	Name   string
	Cards  []Card
	Hand   []Card
	Gems   map[Gem]int
	Nobles []Noble
}

// EffectiveGems returns the combined effectiveness of the cards in a player's possession as well as their gems
func (p *Player) EffectiveGems() map[Gem]int {
	effective := map[Gem]int{}

	for g, v := range p.Gems {
		effective[g] = v
	}

	for _, c := range p.Cards {
		effective[c.Provides]++
	}

	return effective
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

// VictoryPoints returns the vitory points held by a player
func (p *Player) VictoryPoints() int {
	points := 0

	for _, c := range p.Cards {
		points += c.VictoryPoints
	}

	for _, n := range p.Nobles {
		points += n.VictoryPoints
	}

	return points
}

// Deduct the cost of a card from a player's gems
func (p *Player) Deduct(cost map[Gem]int) (map[Gem]int, error) {
	cardGems := map[Gem]int{}
	for _, c := range p.Cards {
		cardGems[c.Provides]++
	}

	handGems := map[Gem]int{}
	for g, c := range p.Gems {
		handGems[g] = c
	}

	payment := map[Gem]int{}

	for gem, count := range cost {
		// subtract gems provided by cards from the cost of the payment required
		count = min(0, count-cardGems[gem])

		if count <= 0 {
			continue
		}

		if count > handGems[gem] {
			return nil, fmt.Errorf("unable to debit player: need %d %s(s), have %d", count, gem.String(), handGems[gem])
		}

		payment[gem] = count
		handGems[gem] -= count
	}

	p.Gems = handGems

	return payment, nil
}

// GiveCard gives a player a card
func (p *Player) GiveCard(card Card) {
	p.Cards = append(p.Cards, card)
}

// ReserveCard gives a player a card in reserve and gems
func (p *Player) ReserveCard(card Card, gems map[Gem]int) {
	p.Hand = append(p.Hand, card)

	p.GiveGems(gems)
}

// GiveGems gives a player gems
func (p *Player) GiveGems(gems map[Gem]int) {
	for g, c := range gems {
		p.Gems[g] += c
	}
}

// GiveNoble gives a player a noble
func (p *Player) GiveNoble(noble Noble) {
	p.Nobles = append(p.Nobles, noble)
}

// Move allows a player to take a turn
func (p *Player) Move(g *Game) {

	// can do one of three things:
	// 1: get gems - either 2 of the same kind or 1 of 3 different kinds - excluding gold
	// 2: reserve card from any tier or top of any deck, receive a gold
	// 3: buy a card on the board or from their

	for moved := false; !moved; {
		switch rand.Intn(3) {
		case 0:
			moved = p.getGems(g)
		case 1:
			moved = p.reserveCard(g)
		case 2:
			moved = p.buyCard(g)
		}
	}

	p.downGems(g, 10)
}

func (p *Player) getGems(g *Game) bool {
	return false
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func (p *Player) reserveCard(g *Game) bool {

	if len(p.Hand) >= 3 {
		return false
	}

	var tier *Deck

	switch rand.Intn(3) {
	case 0:
		tier = &g.Tier1
	case 1:
		tier = &g.Tier2
	case 2:
		tier = &g.Tier3
	}

	if len(*tier) <= 0 {
		return false
	}

	reserveIndex := rand.Intn(min(4, len(*tier)))

	_ = reserveIndex

	return false
}

func (p *Player) buyCard(g *Game) bool {
	return false
}

func (p *Player) downGems(g *Game, max int) {

}
