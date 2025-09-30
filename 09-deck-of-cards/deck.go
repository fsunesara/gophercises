//go:generate stringer -type=Suit

package deck

import (
	"strconv"
)

type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

type Card struct {
	Suit Suit
	Rank int
}

func (c *Card) RankCharacter() string {
	switch c.Rank {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(c.Rank)
	}
}

type deckConfig struct {
	// comparisonFunction func(a, b Card) bool
	// shuffle            bool
	// numJokers          int
	// filterRanks        []int
	// filterSuits        []Suit
	numDecks int
}

func constructDeck() []Card {
	d := []Card{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := 1; rank <= 13; rank++ {
			d = append(d, Card{Suit: suit, Rank: rank})
		}
	}
	return d
}

func New(options ...func(*deckConfig)) *[]Card {
	dc := &deckConfig{}
	for _, option := range options {
		option(dc)
	}

	d := []Card{}
	if dc.numDecks == 0 {
		dc.numDecks = 1
	}
	for i := 0; i < dc.numDecks; i++ {
		d = append(d, constructDeck()...)
	}

	return &d
}

func WithComparisonFunction() {

}

func WithShuffle(shuffle bool) {
}
