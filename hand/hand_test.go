package hand

import (
	"cardgames/card"
	"fmt"
)

func ExampleHand_Add() {
	h := NewHand()
	h.Add(card.New(card.Seven, card.Diamonds))
	h.Add(card.New(card.Eight, card.Clubs))
	h.Add(card.New(card.King, card.Hearts))

	fmt.Printf("%s", h.String())

	// Output:
	// ┌───────┐ ┌───────┐ ┌───────┐
	// │7      │ │8      │ │K      │
	// │       │ │       │ │       │
	// │   ♦   │ │   ♣   │ │   ♥   │
	// │       │ │       │ │       │
	// │      7│ │      8│ │      K│
	// └───────┘ └───────┘ └───────┘
	//
	//    (1)       (2)       (3)
}
