package player

import (
	"cardgames/card"
	"cardgames/hand"
	"fmt"
)

// Human definiert durch Namen und Hand
type Human struct {
	Name string
	Hand hand.Hand
}

// NewHuman() erwartet einen Namen als String und erstellt ein Objekt des Typs Human mit leerer Hand
func NewHuman(name string) *Human {

	return &Human{
		Name: name,
		Hand: hand.NewHand(),
	}
}

// GetName gibt den Namen eines Objekts des Typs Human zurück
func (h Human) GetName() string {
	return h.Name
}

// GetHand gibt die Hand eines Objekts des Typs Human zurück
func (h Human) GetHand() hand.Hand {
	return h.Hand
}

// GetMove fragt den Spieler nach einem Zug. Er gibt die Zahl unter der Karte an. Die Zahl die zurückgegeben wird ist die Position der Karte in seiner Hand
func (h Human) GetMove() int {

	var num int

	fmt.Printf("%s, it´s your turn. Which Card do you want to play?", h.Name)
	fmt.Scan(&num)

	return num - 1
}

// AddCard fügt der Hand des Spielers eine Karte hinzu
func (h *Human) AddCard(c card.Card) {
	h.Hand.Add(c)
}

// RemoveCard löscht eine Karte aus der Hand eines Spielers
func (h *Human) RemoveCard(c card.Card) {
	h.Hand.Remove(c)
}
