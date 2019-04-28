package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	testCases := map[string]struct {
		expected interface{}
		actual   interface{}
	}{
		"Gold": {
			expected: Gem(0),
			actual:   Gold,
		},
		"Diamond": {
			expected: Gem(1),
			actual:   Diamond,
		},
		"Emerald": {
			expected: Gem(2),
			actual:   Emerald,
		},
		"Jet": {
			expected: Gem(3),
			actual:   Jet,
		},
		"Ruby": {
			expected: Gem(4),
			actual:   Ruby,
		},
		"Sapphire": {
			expected: Gem(5),
			actual:   Sapphire,
		},
		"Gold String": {
			expected: "gold",
			actual:   goldStr,
		},
		"Diamond String": {
			expected: "diamond",
			actual:   diamondStr,
		},
		"Emerald String": {
			expected: "emerald",
			actual:   emeraldStr,
		},
		"Jet String": {
			expected: "jet",
			actual:   jetStr,
		},
		"Ruby String": {
			expected: "ruby",
			actual:   rubyStr,
		},
		"Sapphire String": {
			expected: "sapphire",
			actual:   sapphireStr,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.actual)
		})
	}
}

func TestGemMarshalJSON(t *testing.T) {
	quotesAround := func(s string) []byte {
		return []byte(fmt.Sprintf("\"%s\"", s))
	}

	testCases := map[string]struct {
		gem   Gem
		bytes []byte
		err   bool
	}{
		"gold": {
			gem:   Gold,
			bytes: quotesAround(goldStr),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			bytes, err := tc.gem.MarshalJSON()

			assert := assert.New(t)

			if tc.err {
				assert.Error(err)

			} else {
				assert.Equal(tc.bytes, bytes)

				assert.NoError(err)
			}
		})
	}
}

func TestGemUnmarshalJSON(t *testing.T) {
	quotesAround := func(s string) []byte {
		return []byte(fmt.Sprintf("\"%s\"", s))
	}

	testCases := map[string]struct {
		gem   Gem
		bytes []byte
		err   bool
	}{
		"gold": {
			gem:   Gold,
			bytes: quotesAround(goldStr),
		},
		"unknown": {
			bytes: quotesAround("X"),
			err:   true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gem := Gem(0)

			gemPtr := &gem

			err := gemPtr.UnmarshalJSON(tc.bytes)

			assert := assert.New(t)

			if tc.err {
				assert.Error(err)

			} else {
				assert.Equal(tc.gem, gem)

				assert.NoError(err)
			}
		})
	}
}
