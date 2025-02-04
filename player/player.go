package player

import (
	"cardgames/card"
	"cardgames/hand"
)

type Player interface {

	//GetName liefert den Namen des Spielers als String
	GetName() string

	//GetHand liefert die Hand des Spielers als "Hand"
	GetHand() hand.Hand

	//GetMove liefert den Zug des Spielers als int
	GetMove() int

	//AddCard fügt Karten zur Hand eines Spielers hinzu
	AddCard(c card.Card)

	//RemoveCard löscht Karten aus der Hand eines Spielers
	RemoveCard(c card.Card)
}
