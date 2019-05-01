package main

import "fmt"

// Player is the entity playing Luxury
type Player struct {
	Cards []Card
	Hand  []Card
	Gems  map[Gem]int
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
