package dynamic

import (
	. "github.com/sidheart/algorithms/util"
)

// Given an initial hand, a deck, and the index of the next available card in that deck, return the hand that a player
// following the "Stand on 17" strategy would end up with as well as the index of the next card in the deck
// (this index may be out of bounds with regards to the deck)
func standOn17(hand Hand, deck Deck, nextCard int) (Hand, int) {
	for ; nextCard < len(deck.Cards) && hand.BlackjackValue() < 17 && hand.BlackjackValue() != -1; nextCard++ {
		hand.Cards = append(hand.Cards, deck.Cards[nextCard])
	}
	return hand, nextCard
}

// Plays Blackjack optimally against a "Stand on 17" dealer for any given deck of cards
func PerfectBlackjack(deck Deck) int {
	n := len(deck.Cards)
	if n < 4 { // Not enough cards to play a round of Blackjack
		return 0
	}
	memo := make([]int, n, n)
	for i := n - 4; i >= 0; i-- {
		dealerHand := Hand{[]Card{deck.Cards[i], deck.Cards[i+2]}}
		playerHand := Hand{[]Card{deck.Cards[i+1], deck.Cards[i+3]}}
		nextCard := i+4
		var dealerNextCard int
		// Case where player stands
		dealerHand, dealerNextCard = standOn17(dealerHand, deck, nextCard)
		bestOutcome := playerHand.Compare(dealerHand)
		if dealerNextCard < n {
			bestOutcome += memo[dealerNextCard]
		}
		for hits := 0; hits < n - nextCard; hits++ { // All cases where the player hits a valid # of times
			if playerHand.BlackjackValue() == -1 {
				break
			}
			playerHand.Cards = append(playerHand.Cards, deck.Cards[nextCard + hits])
			dealerHand, dealerNextCard = standOn17(dealerHand, deck, nextCard)
			outcome := playerHand.Compare(dealerHand)
			if dealerNextCard < n {
				outcome += memo[dealerNextCard]
			}
			if outcome > bestOutcome {
				bestOutcome = outcome
			}
		}
		memo[i] = bestOutcome
	}
	return memo[0]
}
