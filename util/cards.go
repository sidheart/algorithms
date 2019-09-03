package util

import "math/rand"

type Suite int
type Face int

func (suite Suite) String() string {
	switch suite {
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	}
	panic("Not a valid Suite")
}

func (f Face) String() string {
	switch f {
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	}
	panic("Not a valid Face")
}

func (f Face) value() int {
	switch f {
	case Jack:
		return 10
	case Queen:
		return 10
	case King:
		return 10
	case Ace:
		return 11
	default:
		return int(f)
	}
}

const (
	Diamonds Suite = iota
	Clubs
	Hearts
	Spades
)

const (
	Two Face = 2
	Three = 3
	Four = 4
	Five = 5
	Six = 6
	Seven = 7
	Eight = 8
	Nine = 9
	Ten = 10
	Jack = 11
	Queen = 12
	King = 13
	Ace = 14
)

const Blackjack = 21

var suites = [4]Suite{Diamonds, Clubs, Hearts, Spades}
var faces = [13]Face{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

type Card struct {
	Face Face
	Suite Suite
}

type Deck struct {
	Cards []Card
}

type Hand Deck

func (deck *Deck) shuffle() {
	n := len(deck.Cards)
	for i := 0; i < n; i++ {
		k := rand.Intn(n - i) + i
		temp := deck.Cards[i]
		deck.Cards[i] = deck.Cards[k]
		deck.Cards[k] = temp
	}
}

func StandardDeck() Deck {
	deck := Deck{make([]Card, 52)}
	i := 0
	for _, suite := range suites {
		for _, face := range faces {
			deck.Cards[i] = Card{face, suite}
			i++
		}
	}
	return deck
}

func RandomDeck() Deck {
	deck := StandardDeck()
	deck.shuffle()
	return deck
}

// Returns the total BlackjackValue of a blackjack hand or -1 if the hand is bust
func (hand *Hand) BlackjackValue() int {
	sum := 0
	aces := 0
	for _, card := range hand.Cards {
		if card.Face == Ace {
			aces++
		}
		sum += card.Face.value()
	}
	// Aces can be 1 or 11 in BlackjackValue, adjust the sum so that the player doesn't bust, if possible
	for sum > Blackjack && aces > 0 {
		sum -= 10
		aces--
	}
	// Check for bust
	if sum > Blackjack {
		sum = -1
	}
	return sum
}

func (hand Hand) Compare(other Hand) int {
	a := hand.BlackjackValue()
	b := other.BlackjackValue()
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}
