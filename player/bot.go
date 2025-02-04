package player

import (
	"cardgames/card"
	"cardgames/hand"
	"math/rand"
)

// Bot definiert durch Namen und Hand
type Bot struct {
	Name string
	Hand hand.Hand
}

// NewBot() erwartet einen Namen als String und erstellt ein Objekt des Typs Bot mit leerer Hand
func NewBot(name string) *Bot {

	return &Bot{
		Name: name,
		Hand: hand.NewHand(),
	}
}

// GetName gibt den Namen eines Objekts des Typs Bot zurück
func (b Bot) GetName() string {
	return b.Name
}

// GetHand gibt die Hand eines Objekts des Typs Bot zurück
func (b Bot) GetHand() hand.Hand {
	return b.Hand
}

// GetMove gibt einen regelconformen Move zurück
func (b Bot) GetMove(n []int) int {

	if len(n) == 0 {
		return 0
	}

	return rand.Intn(len(n)) + 1
}

// AddCard fügt der Hand des Spielers eine Karte hinzu
func (b *Bot) AddCard(c card.Card) {
	b.Hand.Add(c)
}

// RemoveCard löscht eine Karte aus der Hand eines Spielers
func (b *Bot) RemoveCard(c card.Card) {
	b.Hand.Remove(c)
}
