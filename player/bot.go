package player

import (
	"cardgames/card"
	"cardgames/hand"
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
		Hand: hand.New(),
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
func (b Bot) GetMove() int {

	return 0
}

// AddCard fügt der Hand des Spielers eine Karte hinzu
func (b *Bot) AddCard(c card.Card) {
	b.Hand.Add(c)
}

// RemoveCard löscht eine Karte aus der Hand eines Spielers
func (b *Bot) RemoveCard(c card.Card) {
	b.Hand.Remove(c)
}
