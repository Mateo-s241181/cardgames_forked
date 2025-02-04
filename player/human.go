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

// GetMove fragt den Spieler nach einem Zug. Er gibt die Zahl unter der Karte an. Die Zahl die zurückgegeben wird ist die Position der Karte in seiner Hand, es wird geckecked ob der Move in n enthalten ist.
func (h Human) GetMove(n []int) int {

	var move int

	//Move einlesen
	fmt.Printf("%s, it´s your turn. Which Card do you want to play?\nPress '0' to draw a Card\n", h.Name)
	fmt.Scan(&move)

	//Move in variable speichern und contains zunächst auf false setzen
	contains := false

	if move > len(h.GetHand().Cards) {
		fmt.Println("Your move was out of bounds, try again.")
		return h.GetMove(n)
	}

	//Checken ob move in l ist
	for i := range n {

		if move == n[i] {
			contains = true
		}
	}

	//Wenn Move nicht in Legalmoves drin ist, soll Funktion nochmal ausgeführt werden
	if !contains && move != 0 {
		fmt.Println("Your move was illegal, try again.")
		return h.GetMove(n)
	}

	return move
}

// AddCard fügt der Hand des Spielers eine Karte hinzu
func (h *Human) AddCard(c card.Card) {
	h.Hand.Add(c)
}

// RemoveCard löscht eine Karte aus der Hand eines Spielers
func (h *Human) RemoveCard(c card.Card) {
	h.Hand.Remove(c)
}
