package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Suit: Joker})
}

func TestNew(t *testing.T) {
	cards := New(DefaultSort)

	if len(cards) != 13*4 {
		t.Errorf("Wrong number of card")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Errorf("Expected Ace and Spade for Card 0")
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Errorf("Expected Ace and Spade for Card 0")
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("Expected three joker")
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("Expected two and three to be filter out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d card, received %d card and deck %d", 13*4*3, len(cards), 3)
	}
}

func TestShuffle(t *testing.T) {
	// Stub value
	shuffleRand = rand.New(rand.NewSource(0))
	// As shuffleRand.Perm(52) gives same result everytime
	// [40 35 50 0 44 7 1 16 13 4 21 12 23 34 19 11 42 20 17 48 27 9 43 46 47 45 5 49 51 30 41 26 25 32 39 28 37 31 33 10 22 8 6 29 36 18 14 2 15 3 38 24]

	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("First shuffle value expected to be %d and got %d", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Second shuffle value expected to be %d and got %d", second, cards[1])
	}
}
