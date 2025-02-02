package card

import (
	"fmt"
	"strings"
)

type Card struct {
	S Suit
	R Rank
}

func New(r Rank, s Suit) Card {
	return Card{
		S: s,
		R: r,
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
	card_string := fmt.Sprintf(card_string_template, c.R, c.S, c.R)
	return card_string
}

// Matches erwartet eine weitere Karte o und prüft,
// ob c und o nach den MauMau-Regeln zusammenpassen.
func (c Card) Matches(o Card) bool {
	// Wenn Farbe oder Zahl gleich sind, passen die
	// Karten zusammen
	return c.R == o.R || c.S == o.S
}

// GetRank gibt den Rang einer Karte zurück
func GetRank(c Card) Rank {
	return c.R
}

// GetSuit gibt den Suit einer Karte zurück
func GetSuit(c Card) Suit {
	return c.S
}
