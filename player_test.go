package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGiveCard(t *testing.T) {
	testCases := map[string]struct {
		player   *Player
		card     Card
		expected []Card
	}{
		"no cards": {
			player:   &Player{},
			card:     Card{},
			expected: []Card{{}},
		},
		"add card": {
			player:   &Player{Cards: []Card{{}}},
			card:     Card{},
			expected: []Card{{}, {}},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tc.player.GiveCard(tc.card)

			assert.Equal(t, tc.expected, tc.player.Cards)
		})
	}
}

func TestGiveGems(t *testing.T) {
	type gemCount struct {
		gem   Gem
		count int
	}

	makeGems := func(gcs ...gemCount) map[Gem]int {
		out := map[Gem]int{}
		for _, gc := range gcs {
			out[gc.gem] += gc.count
		}

		return out
	}

	removeZeroKeys := func(gems map[Gem]int) {
		for k, v := range gems {
			if v == 0 {
				delete(gems, k)
			}
		}
	}

	testCases := map[string]struct {
		player   *Player
		gems     map[Gem]int
		expected map[Gem]int
	}{
		"new gems": {
			player:   &Player{Gems: makeGems()},
			gems:     makeGems(gemCount{Emerald, 1}),
			expected: makeGems(gemCount{Emerald, 1}),
		},
		"different gems": {
			player:   &Player{Gems: makeGems(gemCount{Jet, 1})},
			gems:     makeGems(gemCount{Emerald, 1}),
			expected: makeGems([]gemCount{{Emerald, 1}, {Jet, 1}}...),
		},
		"existing gems": {
			player:   &Player{Gems: makeGems(gemCount{Emerald, 1})},
			gems:     makeGems(gemCount{Emerald, 2}),
			expected: makeGems(gemCount{Emerald, 3}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tc.player.GiveGems(tc.gems)

			removeZeroKeys(tc.player.Gems)
			removeZeroKeys(tc.expected)

			assert.Equal(t, tc.expected, tc.player.Gems)
		})
	}
}
