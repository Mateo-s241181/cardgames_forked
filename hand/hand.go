package hand

import (
	"cardgames/card"
	"fmt"
)

type Hand struct {
	Cards []card.Card
}

// New gibt eine leere Hand zurück.
func NewHand() Hand {
	return Hand{}
}

// Add fügt eine Karte zur Hand hinzu.
func (h *Hand) Add(c card.Card) {
	h.Cards = append(h.Cards, c)
}

// String gibt eine AsciiArt-Repräsentation der Hand zurück.
func (h Hand) String() string {

	//TODO: Alle Print Statements so umwandeln, dass die in einen String geschrieben Werden, der ausgegeben werden kann.
	s := ""

	//Erste Zeile
	addStringNTimes(&s, "┌───────┐ ", h)

	//Zweite Zeile
	for i := range h.Cards {
		s += fmt.Sprintf("│%-2s     │ ", card.GetRank(h.Cards[i]))
	}

	s += "\n"

	//Dritte Zeile
	addStringNTimes(&s, "│       │ ", h)

	//Vierte Zeile
	for i := range h.Cards {
		s += fmt.Sprintf("│   %s   │ ", card.GetSuit(h.Cards[i]))
	}

	s += "\n"

	//Fünfte Zeile
	addStringNTimes(&s, "│       │ ", h)

	//Sechste Zeile
	for i := range h.Cards {
		s += fmt.Sprintf("│     %2s│ ", card.GetRank(h.Cards[i]))
	}

	s += "\n"

	//Siebte Zeile
	addStringNTimes(&s, "└───────┘ ", h)

	s += "\n"

	for i := range h.Cards {
		s += fmt.Sprintf("   (%d)    ", i+1)
	}

	return s
}

// Hilfsfunktion um einen String n mal auszugeben
func addStringNTimes(s *string, addString string, h Hand) {

	//Addstring n mal an s anhängen
	for i := 0; i < len(h.Cards); i++ {
		*s += addString
	}

	*s += "\n"
}

// Remove entfernt eine Karte aus der Hand.
func (h *Hand) Remove(c card.Card) {

	//Durch h.Cards rangen
	for i := range h.Cards {
		//Wenn c an der Stelle i ist
		if h.Cards[i] == c {

			//h.Cards[i] löschen

			//if conditions um Index + 1 in Range zu behalten
			if i != len(h.Cards)-1 {
				h.Cards = append(h.Cards[:i], h.Cards[i+1:]...)
			} else {
				h.Cards = h.Cards[:i]
			}
			break
		}
	}
}

// Len gibt die Anzahl der Karten in der Hand zurück.
func (h Hand) Len() int {
	return len(h.Cards)
}

// ContainsRank prüft, ob ein bestimmter Rang in der Hand ist.
func (h Hand) ContainsRank(r card.Rank) bool {

	for i := range h.Cards {

		if card.GetRank(h.Cards[i]) == r {
			return true
		}
	}

	return false
}
