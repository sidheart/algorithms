package dynamic

import (
	"fmt"
	. "github.com/sidheart/algorithms/util"
)

// Represents what will happen during a single round of Blackjack
type Round struct {
	numHits int
	nextCard int
}

type DealerStrategy func(Hand, Deck, int) (Hand, int)

// Given an initial hand, a deck, and the index of the next available card in that deck, return the hand that a player
// following the "Stand on 17" strategy would end up with as well as the index of the next card in the deck
// (this index may be out of bounds with regards to the deck)
func standOn17(hand Hand, deck Deck, nextCard int) (Hand, int) {
	for ; nextCard < len(deck.Cards) && hand.BlackjackValue() < 17 && hand.BlackjackValue() != -1; nextCard++ {
		hand.Cards = append(hand.Cards, deck.Cards[nextCard])
	}
	return hand, nextCard
}

func printRounds(optimalRounds []Round, totalCards int) {
	nextCard := 0
	turn := 1
	for nextCard < totalCards {
		fmt.Printf("On turn %d you should: ", turn)
		r := optimalRounds[nextCard]
		if r.numHits == 0 {
			fmt.Printf("stand\n")
		} else {
			fmt.Printf("hit %d times\n", r.numHits)
		}
		nextCard = r.nextCard
		turn++
	}
	fmt.Println()
}

// Plays Blackjack optimally against a "Stand on 17" dealer for any given deck of cards
func PerfectBlackjack(deck Deck) int {
	n := len(deck.Cards)
	if n < 4 { // Not enough cards to play a round of Blackjack
		return 0
	}
	memo := make([]int, n, n)
	parent := make([]Round, n, n)
	for i := n - 4; i >= 0; i-- {
		// Assume the dealer always deals to himself first, then the player, then himself, then the player one last time
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
		parent[i] = Round{numHits: 0, nextCard: dealerNextCard}
		// All cases where the player hits a valid # of times, also note that "0" means hit once, "1" is twice, etc.
		for hits := 0; hits < n - nextCard; hits++ {
			// Note: even if the hand value is 21, the player might want to intentionally bust to improve future hands
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
				parent[i] = Round{numHits: hits + 1, nextCard: dealerNextCard}
			}
		}
		memo[i] = bestOutcome
	}
	printRounds(parent, n)
	return memo[0]
}
