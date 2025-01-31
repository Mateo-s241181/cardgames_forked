package hand

import "cardgames/card"

type Hand struct {
	Cards []card.Card
}

// New gibt eine leere Hand zurück.
func New() Hand {
	return Hand{}
}

// Add fügt eine Karte zur Hand hinzu.
func (h *Hand) Add(c card.Card) {
	h.Cards = append(h.Cards, c)
}

// String gibt eine AsciiArt-Repräsentation der Hand zurück.
func (h Hand) String() string {
	// TODO
	return ""
}

// Remove entfernt eine Karte aus der Hand.
func (h *Hand) Remove(c card.Card) {
	// TODO
}

// Len gibt die Anzahl der Karten in der Hand zurück.
func (h Hand) Len() int {
	// TODO
	return 0
}

// ContainsRank prüft, ob ein bestimmter Rang in der Hand ist.
func (h Hand) ContainsRank(r card.Rank) bool {
	// TODO
	return false
}
