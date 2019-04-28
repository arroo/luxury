package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapper(t *testing.T) {
	deck := Deck{}

	fn := deck.Swapper()

	assert.NotNil(t, fn)
}
