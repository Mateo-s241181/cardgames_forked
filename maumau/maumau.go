package main

import (
	"cardgames/game"
	"fmt"
)

func main() {

	winner := -1

	gameEnds := false
	counter := 0

	//Spiel mit leeren Spielerstrukturen erstellen
	g := game.NewGame(game.PlayerCount())

	//Playerstrukturen in Game füllen
	g.CreatePlayer()

	//Karten mischen und an die Spieler verteilen
	g.Deck.Shuffle()
	g.DistributeCards()

	//Oberste Karte vom Deck entfernen und auf den Ablagestapel legen
	g.FriedhofInit()

	for !gameEnds {
		fmt.Println("This is the Card on top of the deck:")
		fmt.Println(g.Friedhof.Cards[0])

		//Counter resetten
		if counter%len(g.Players) == 0 {
			counter = 0
		}

		//l speichert die legalen Moves, die ein Spieler auf die TopCard vom friedhof legen kann
		l := g.LegalMoves(g.Friedhof.Top(), counter)

		//Hand printen
		fmt.Println("This is your hand:")
		fmt.Println(g.Players[counter].GetHand().String())

		//Move von Spieler einlesen
		moveAsInt := g.Players[counter].GetMove(l)

		//fmt.Printf("Move: %d\n", moveAsInt)
		//fmt.Printf("Legal Moves: %v", l)

		//Move ausführen
		g.Move(moveAsInt, counter)

		//Checken ob Spieler noch Karten hat
		if len(g.Players[counter].GetHand().Cards) == 0 {

			//Counter startet bei null, daher plus 1
			winner = counter

			gameEnds = true
		}

		//counter erhöhen
		counter++
	}

	fmt.Println("-------------------------------")
	fmt.Printf("       %s wins!         \n", g.Players[winner].GetName())
	fmt.Println("-------------------------------")
}

//Wie viele Spieler gibt es?

/*
Ablauf des Spiels:

(1) Initialisieren des Spiels

		- Kartendeck erstellen
		- Wie viele Spieler gibt es?
		- Wie heißen die Spieler und sind es Bots?

	(2) Vorkehrungen vor Spielstart

		- Karten werden gemischt
		- Karten werden in die Hand der Spieler ausgegeben
		- Spielerreihenfolge festgelegt
		- Erste Karte wird aufgedeckt und angezeigt

	(3) Ablauf eines Zuges

		- Auslesen der Obersten Karte
		- Ermitteln der legalen Züge => Abgleichen von Rank und Suit der Handkarten

		- Case: Spieler
			- Spieler nach Zug fragen
			- Zug mit legalen Zügen vergleichen
			- Kein Zug möglich => Karte vom Stapel ziehen

		- Case: Bot
			- Zufälliger legaler Zug
			- Kein Zug möglich => Karte vom Stapel Ziehen

		- Karte auf den Ablagestapel legen

	(4) Ende des Spiels

		- Leere Hand => Gewinner

*/
