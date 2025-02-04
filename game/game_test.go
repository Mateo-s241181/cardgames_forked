package game

import (
	"cardgames/card"
	"cardgames/player"
	"fmt"
)

func ExampleGame_PlayerCount() {

	var g Game

	p := g.PlayerCount()

	fmt.Println(p)

	// Output:
	// Kann nicht getestet werden, da Funktion PlayerCount einen Input erwartet
}

func Example_NewGame() {

	fmt.Println(NewGame(4))

	//Output:
	//
}

func ExampleGame_DistributeCards() {

	g := NewGame(4)

	g.Deck.Shuffle()

	for i := range g.Players {
		g.Players[i] = player.NewBot(fmt.Sprintf("Player %d", i+1))
	}

	g.DistributeCards()

	g.Players[0].AddCard(card.New(card.Ace, card.Spades))

	for i := range g.Players {
		fmt.Println(g.Players[i].GetName())
		fmt.Println(g.Players[i].GetHand().String())

		fmt.Println(g.LegalMoves(card.New(card.Ace, card.Clubs), 3))
	}
	// Output:
	// 1
}

func ExampleGame_LegalMoves() {

	g := NewGame(3)

	for i := range g.Players {
		g.Players[i] = player.NewBot(fmt.Sprintf("Player %d", i+1))
	}

	top := card.New(card.Ace, card.Diamonds)

	c := []card.Card{

		card.New(card.Ace, card.Clubs),
		card.New(card.Seven, card.Spades),
		card.New(card.Eight, card.Spades),
		card.New(card.Nine, card.Spades),
	}

	for i := range c {
		g.Players[0].AddCard(c[i])
	}

	fmt.Println(g.LegalMoves(top, 0))

	// Output:
}
