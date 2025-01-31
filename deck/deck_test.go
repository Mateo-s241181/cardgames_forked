package deck

import (
	"cardgames/card"
	"fmt"
)

func ExampleDeck_new() {
	d := NewDeck32()
	fmt.Println(d.Len())

	// Output:
	// 32
}

func ExampleDeck_Draw() {
	d := NewDeck32()
	c := d.Draw()
	fmt.Println(c)
	fmt.Println(d.Len())

	// Output:
	// ┌───────┐
	// │7      │
	// │       │
	// │   ♣   │
	// │       │
	// │      7│
	// └───────┘
	// 31
}

func ExampleDeck_Top() {
	d := NewDeck32()
	c := d.Top()
	fmt.Println(c)
	fmt.Println(d.Len())

	// Output:
	// ┌───────┐
	// │7      │
	// │       │
	// │   ♣   │
	// │       │
	// │      7│
	// └───────┘
	// 32
}

func ExampleDeck_Add() {
	d := NewDeck32()
	c := card.New(card.Seven, card.Diamonds)
	d.Add(c)
	fmt.Println(d.Len())
	fmt.Println(d.Top())

	// Output:
	// 33
	// ┌───────┐
	// │7      │
	// │       │
	// │   ♦   │
	// │       │
	// │      7│
	// └───────┘
}

func ExampleDeck_Shuffle() {
	d := NewDeck32()

	d.Shuffle()

	fmt.Println(d)
	fmt.Println(len(d.Cards))

	// Output:
	// foo
}
