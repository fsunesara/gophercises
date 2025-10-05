package deck

import "testing"

func TestNumJokers(t *testing.T) {
	got := New(WithNumJokers(0))
	if len(got) != 52 {
		t.Errorf("expected 52 cards, got %d", len(got))
	}

	got = New(WithNumJokers(1))
	if len(got) != 53 {
		t.Errorf("expected 53 cards, got %d", len(got))
	}
}

func TestFilterRanks(t *testing.T) {
	got := New(WithFilterRanks([]int{1, 2, 3}))
	if len(got) != 40 {
		t.Errorf("expected 40 cards, got %d", len(got))
	}
}

func TestFilterSuits(t *testing.T) {
	got := New(WithFilterSuits([]Suit{Spades, Hearts}))
	if len(got) != 26 {
		t.Errorf("expected 26 cards, got %d", len(got))
	}
}

func TestNumDecks(t *testing.T) {
	got := New(WithNumDecks(0))
	if len(got) != 52 {
		t.Errorf("expected 52 cards, got %d", len(got))
	}

	got = New(WithNumDecks(1))
	if len(got) != 52 {
		t.Errorf("expected 52 cards, got %d", len(got))
	}

	got = New(WithNumDecks(2))
	if len(got) != 104 {
		t.Errorf("expected 104 cards, got %d", len(got))
	}
}
