package dynamic

import (
	"fmt"
	. "github.com/sidheart/algorithms/util"
	"math"
)

type DealerStrategy func(Hand, Hand, Deck, int) (Hand, int)

type Outcome struct {
	totalValue int
	hits int
}

// Given an initial hand, a deck, and the index of the next available card in that deck, return the hand that a player
// following the "Stand on 17" strategy would end up with as well as the index of the next card in the deck
// (this index may be out of bounds with regards to the deck)
func standOn17(dealerHand Hand, playerHand Hand, deck Deck, nextCard int) (Hand, int) {
	// Dealer won't hit if the player is bust in this particular strategy
	if playerHand.BlackjackValue() == -1 {
		return  dealerHand, nextCard
	}
	for ; nextCard < len(deck.Cards) && dealerHand.BlackjackValue() <= playerHand.BlackjackValue() && dealerHand.BlackjackValue() < 17 && dealerHand.BlackjackValue() != -1; nextCard++ {
		dealerHand.Cards = append(dealerHand.Cards, deck.Cards[nextCard])
	}
	return dealerHand, nextCard
}

func printStrategy(solution map[int]Outcome, n int) {
	for i, round := 0, 1; i < n - 3; {
		bestOutcome := solution[i]
		fmt.Printf("On round %d you should ", round)
		if bestOutcome.hits == 0 {
			fmt.Printf("stand.\n")
		} else {
			fmt.Printf("hit %d times.\n", bestOutcome.hits)
		}
		i += bestOutcome.hits + 4
		round++
	}
	fmt.Println()
}

func PerfectBlackjack(deck Deck) int {
	solution := make(map[int]Outcome)
	perfectBlackjackTopdown(deck, 0, standOn17, solution)
	printStrategy(solution, len(deck.Cards))
	return solution[0].totalValue
}

func perfectBlackjackTopdown(deck Deck, i int, dealerStrategy DealerStrategy, memo map[int]Outcome) int {
	n := len(deck.Cards)
	if n - i < 4 {
		memo[i] = Outcome{0, 0}
		return 0
	}
	dealerHand := Hand{[]Card{deck.Cards[i], deck.Cards[i+2]}}
	playerHand := Hand{[]Card{deck.Cards[i+1], deck.Cards[i+3]}}
	busted := false
	bestOutcome := math.MinInt64
	for j := i+3; !busted && j < n; j++ {
		if j != i + 3 { // When j == i+3, the player stands
			playerHand.Cards = append(playerHand.Cards, deck.Cards[j])
		}
		if playerHand.BlackjackValue() > Blackjack {
			busted = true
		}
		dealerHand, nextCard := dealerStrategy(dealerHand, playerHand, deck, j+1)
		outcome := playerHand.Compare(dealerHand)
		nextBestOutcome, ok := memo[nextCard]
		if ok {
			outcome += nextBestOutcome.totalValue
		} else {
			outcome += perfectBlackjackTopdown(deck, nextCard, dealerStrategy, memo)
		}
		if outcome > bestOutcome {
			bestOutcome = outcome
			memo[i] = Outcome{bestOutcome, j - i - 3}
		}
	}
	return memo[i].totalValue
}
