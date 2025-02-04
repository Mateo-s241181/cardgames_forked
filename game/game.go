package game

import (
	"cardgames/card"
	"cardgames/deck"
	"cardgames/player"
	"fmt"
)

//Start des Spiels:

//Man Spielt mit 3 Karten

var NUMBER_OF_CARDS int = 3

// Das Spiel wird Definiert durch ein Deck und eine Liste an Spielern
type Game struct {
	Deck     deck.Deck
	Friedhof deck.Deck
	Players  []player.Player
}

//---------------------------------------------------------------------------------
//(1) Spiel intinialisieren
//---------------------------------------------------------------------------------

// Fragt nach Input die Anzahl der Spieler betreffend
func PlayerCount() int {

	var playerCount int

	fmt.Println("How many Players play the Game? (2-4)")

	fmt.Scan(&playerCount)

	//Falls Input nicht zwischen 2 und 4 liegt nochmal abfragen
	if playerCount < 2 || playerCount > 4 {
		fmt.Println("False Input")
		return PlayerCount()
	}

	return playerCount
}

// NewGame erwartet eine Anzahl von Spielern und gibt ein Objekt des Typs Game mit einem 32 Karten Deck und einer leeren Liste an Spielern der Länge p aus
func NewGame(p int) Game {

	return Game{
		Deck:     deck.NewDeck32(),
		Friedhof: deck.NewEmptyDeck(),
		Players:  make([]player.Player, p),
	}
}

// CreatePlayer fragt nach den Namen der einzelnen Spieler
// Ändert die Namen in der Player Liste und gibt das Game zurück

func (g Game) CreatePlayer() {

	for i := range g.Players {

		var name string
		var isBot string

		//Namen abfragen
		fmt.Printf("Enter a name: ")
		fmt.Scan(&name)
		fmt.Printf("\n")

		//isBot abfragen
		fmt.Printf("Is Player %d a Bot? (y/n) \n", i+1)
		fmt.Scan(&isBot)
		fmt.Printf("\n")

		//Erstellung eines Spielers abhängig von isBot
		if isBot == "y" {
			g.Players[i] = player.NewBot(name)
		}
		if isBot == "n" {
			g.Players[i] = player.NewHuman(name)
		}
		if isBot != "y" && isBot != "n" {
			fmt.Println("Falsche Eingabe")

			//Schleife soll die Eingabe wiederholen
			i--
			continue
		}

	}
}

//-----------------------------------------------------------------------------
//(2) Vorkehrungen vor Spielstart
//-----------------------------------------------------------------------------

//Mischfunktion bereits in Deck definiert

// Karten in die Hand des jeweiligen Spielers ausgeben
// Funktion hierzu muss in Interface Player enthalten sein

func (g *Game) DistributeCards() {

	for i := range g.Players {

		for j := 0; j < NUMBER_OF_CARDS; j++ {
			g.Players[i].AddCard(g.Deck.Draw())
		}
	}
}

//Spielerreihenfolge durch Reihenfolge in g.Players festgelegt

//Ablagestapel als Deck definieren (in Deck funktion)

func (g *Game) FriedhofInit() {

	top := g.Deck.Draw()

	g.Friedhof.Add(top)
}

//------------------------------------------------------------------------------
//(3) Ablauf eines Zuges
//------------------------------------------------------------------------------

//Auslesen der obersten Karte in Deck durch Top definiert
//=> Oberste Karte muss auf den Ablagestapel gelegt werden

// Ermitteln der Legalen Züge in Slice
// Nimmt als Argument die oberste Karte und den Spieler der im Moment am Zug ist
func (g Game) LegalMoves(c card.Card, n int) []int {

	legal := []int{}

	h := g.Players[n].GetHand().Cards

	for i := range h {

		if h[i].Matches(c) {
			legal = append(legal, i+1)
		}
	}

	return legal
}

//Spieler Abfragen

// Zug ausführen => Karte auf Ablagestabe legen oder ziehen
// Input == Spieler, der an der Reihe ist
func (g *Game) Move(move, playerNumber int) {

	if move == 0 {
		//Spieler zieht eine Karte vom Stapel
		g.Players[playerNumber].AddCard(g.Deck.Draw())
	}

	//c ist die Karte, die der Spieler ablegen möchte
	if move != 0 {
		c := g.Players[playerNumber].GetHand().Cards[move-1]

		//Karte wird an den Anfang des Ablagestapels gelegt, daher die Reihenfolge im append
		g.Friedhof.Cards = append([]card.Card{c}, g.Friedhof.Cards...)

		//Karte wird aus der Hand des Spielers entfernt
		g.Players[playerNumber].RemoveCard(c)
	}
}

//-------------------------------------------------------------------------------
//(4) Ende des Spiels
//-------------------------------------------------------------------------------

//Auslesen der Länge der Handarrays (definiert in Hand)
