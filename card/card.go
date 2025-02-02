package card

import (
	"fmt"
	"strings"
)

type Card struct {
	s Suit
	r Rank
}

func New(r Rank, s Suit) Card {
	return Card{
		s: s,
		r: r,
	}
}

// String liefert eine AsciiArt-Repräsentation der Karte.
// Zeichnet einen Rahmen um die Karte und innen die Zeichen für
// Rang und Farbe.
// Beispiel Pik Ass:
// ┌───────┐
// │A      │
// │       │
// │   ♠   │
// │       │
// │      A│
// └───────┘
func (c Card) String() string {
	lines := []string{
		"┌───────┐",
		"│%-2s     │",
		"│       │",
		"│   %s   │",
		"│       │",
		"│     %2s│",
		"└───────┘",
	}

	card_string_template := strings.Join(lines, "\n")
	card_string := fmt.Sprintf(card_string_template, c.r, c.s, c.r)
	return card_string
}

// Matches erwartet eine weitere Karte o und prüft,
// ob c und o nach den MauMau-Regeln zusammenpassen.
func (c Card) Matches(o Card) bool {
	// Wenn Farbe oder Zahl gleich sind, passen die
	// Karten zusammen
	return c.r == o.r || c.s == o.s
}

func GetRank(c Card) Rank {
	return c.r
}

func GetSuit(c Card) Suit {
	return c.s
}
