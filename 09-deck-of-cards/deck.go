package deck

import (
	"math/rand/v2"
	"slices"
	"strconv"
)

type Suit int

const (
	Spades Suit = iota
	Diamonds
	Hearts
	Clubs
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "♠"
	case Diamonds:
		return "♢"
	case Hearts:
		return "♡"
	case Clubs:
		return "♣"
	default:
		return ""
	}
}

type Card struct {
	Suit *Suit
	Rank int
}

func (c Card) String() string {
	return c.RankToString() + c.Suit.String()
}

func (c *Card) RankToString() string {
	switch c.Rank {
	case 0:
		return "JKR"
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
	comparisonFunction func(a, b Card) int
	shuffle            bool
	numJokers          int
	filterRanks        []int
	filterSuits        []Suit
	numDecks           int
}

func constructDeck(dc *deckConfig) []Card {
	d := []Card{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		if len(dc.filterSuits) > 0 && slices.Contains(dc.filterSuits, suit) {
			continue
		}
		for rank := 1; rank <= 13; rank++ {
			if len(dc.filterRanks) > 0 && slices.Contains(dc.filterRanks, rank) {
				continue
			}
			d = append(d, Card{Suit: &suit, Rank: rank})
		}
	}
	return d
}

func defaultComparisonFunction(a, b Card) int {
	if a.Rank == b.Rank && a.Suit != nil && b.Suit != nil {
		return int(*a.Suit) - int(*b.Suit)
	}
	return a.Rank - b.Rank
}

var comparisonFunction = defaultComparisonFunction

func New(options ...func(*deckConfig)) []Card {
	dc := &deckConfig{}
	for _, option := range options {
		option(dc)
	}

	d := []Card{}
	if dc.numDecks <= 0 {
		dc.numDecks = 1
	}
	for i := 0; i < dc.numDecks; i++ {
		d = append(d, constructDeck(dc)...)
	}

	if dc.numJokers > 0 {
		for i := 0; i < dc.numJokers; i++ {
			d = append(d, Card{Suit: nil, Rank: 0})
		}
	}

	if dc.comparisonFunction != nil {
		comparisonFunction = dc.comparisonFunction
	}

	if dc.shuffle {
		rand.Shuffle(len(d), func(i, j int) {
			d[i], d[j] = d[j], d[i]
		})
	}
	return d
}

func Sort(d []Card) {
	slices.SortFunc(d, comparisonFunction)
}

func WithComparisonFunction(comparisonFunction func(a, b Card) int) func(*deckConfig) {
	return func(dc *deckConfig) {
		dc.comparisonFunction = comparisonFunction
	}
}

func WithShuffle(shuffle bool) func(*deckConfig) {
	return func(dc *deckConfig) {
		dc.shuffle = shuffle
	}
}

func WithNumJokers(numJokers int) func(*deckConfig) {
	return func(dc *deckConfig) {
		dc.numJokers = numJokers
	}
}

func WithFilterRanks(filterRanks []int) func(*deckConfig) {
	return func(dc *deckConfig) {
		dc.filterRanks = filterRanks
	}
}

func WithFilterSuits(filterSuits []Suit) func(*deckConfig) {
	return func(dc *deckConfig) {
		dc.filterSuits = filterSuits
	}
}

func WithNumDecks(numDecks int) func(*deckConfig) {
	return func(dc *deckConfig) {
		dc.numDecks = numDecks
	}
}
