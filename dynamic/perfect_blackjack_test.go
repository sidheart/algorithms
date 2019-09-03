package dynamic

import (
	. "github.com/sidheart/algorithms/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerfectBlackjack(t *testing.T) {
	// Dealer blackjack
	deck := Deck{Cards: []Card{
		{Face: King, Suite: Hearts},
		{Face: Two, Suite: Clubs},
		{Face: Ace, Suite: Hearts},
		{Face: Three, Suite: Clubs},
	}}
	assert.Equal(t, -1, PerfectBlackjack(deck))
	// Basic tie
	deck = Deck{Cards: []Card{
		{Face: King, Suite: Hearts},
		{Face: King, Suite: Clubs},
		{Face: Ace, Suite: Hearts},
		{Face: Ace, Suite: Clubs},
	}}
	assert.Equal(t, 0, PerfectBlackjack(deck))
	// Player blackjack
	deck = Deck{Cards: []Card{
		{Face: Two, Suite: Hearts},
		{Face: King, Suite: Clubs},
		{Face: Ace, Suite: Hearts},
		{Face: Ace, Suite: Clubs},
	}}
	assert.Equal(t, 1, PerfectBlackjack(deck))
	// Player stands on both hands and wins
	deck = Deck{Cards: []Card{
		{Face: Seven, Suite: Hearts},
		{Face: King, Suite: Clubs},
		{Face: Ace, Suite: Hearts},
		{Face: Ace, Suite: Clubs},
		{Face: Two, Suite: Hearts},
		{Face: King, Suite: Clubs},
		{Face: Ace, Suite: Hearts},
		{Face: Ace, Suite: Clubs},
	}}
	assert.Equal(t, 2, PerfectBlackjack(deck))
}
