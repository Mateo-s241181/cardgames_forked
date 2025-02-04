package game

import (
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

	for i := range g.Players {
		fmt.Println(g.Players[i].GetName())
		fmt.Println(g.Players[i].GetHand().String())
	}
	// Output:
	// 1
}
