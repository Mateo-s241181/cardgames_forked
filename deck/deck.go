package deck

import (
	"cardgames/card"
	"math/rand"
)

// Ein Kartenstapel ist eine Liste von Karten.
type Deck struct {
	Cards []card.Card
}

// NewDeck32 gibt einen neuen Kartenstapel zurück.
func NewDeck32() Deck {
	// Skat-Blatt: 32 Karten, 7-10, Bube, Dame, König, Ass
	cards := []card.Card{
		card.New(card.Seven, card.Clubs),
		card.New(card.Eight, card.Clubs),
		card.New(card.Nine, card.Clubs),
		card.New(card.Ten, card.Clubs),
		card.New(card.Jack, card.Clubs),
		card.New(card.Queen, card.Clubs),
		card.New(card.King, card.Clubs),
		card.New(card.Ace, card.Clubs),
		card.New(card.Seven, card.Spades),
		card.New(card.Eight, card.Spades),
		card.New(card.Nine, card.Spades),
		card.New(card.Ten, card.Spades),
		card.New(card.Jack, card.Spades),
		card.New(card.Queen, card.Spades),
		card.New(card.King, card.Spades),
		card.New(card.Ace, card.Spades),
		card.New(card.Seven, card.Hearts),
		card.New(card.Eight, card.Hearts),
		card.New(card.Nine, card.Hearts),
		card.New(card.Ten, card.Hearts),
		card.New(card.Jack, card.Hearts),
		card.New(card.Queen, card.Hearts),
		card.New(card.King, card.Hearts),
		card.New(card.Ace, card.Hearts),
		card.New(card.Seven, card.Diamonds),
		card.New(card.Eight, card.Diamonds),
		card.New(card.Nine, card.Diamonds),
		card.New(card.Ten, card.Diamonds),
		card.New(card.Jack, card.Diamonds),
		card.New(card.Queen, card.Diamonds),
		card.New(card.King, card.Diamonds),
		card.New(card.Ace, card.Diamonds),
	}

	return Deck{Cards: cards}
}

// Mischt das Deck.
func (d *Deck) Shuffle() {

	//Kreiert ein neue cards Liste "Shuffled"

	shuffled := []card.Card{}

	//Wiederholen solange d.Cards noch Elemente enthält
	for len(d.Cards) != 0 {

		random := rand.Intn(len(d.Cards))
		//Zieht eine random Karte aus d.Cards und appendet sie an "Shuffled"
		shuffled = append(shuffled, d.Cards[random])

		//Löscht die gezogene Karte aus d.Cards
		d.Cards = append(d.Cards[:random], d.Cards[random+1:]...)

		//Fortfahren bis d.Cards leer ist
	}

	//Am Ende d.Cards = "Shuffled"
	d.Cards = shuffled
}

// Zieht eine Karte vom Deck.
// D.h. entfernt die oberste Karte und gibt sie zurück.
func (d *Deck) Draw() card.Card {

	//Oberste Karte auslesen
	c := d.Cards[0]

	//Oberste Karte aus dem Deck entfernen
	d.Cards = d.Cards[1:]

	return c
}

// Gibt die oberste Karte des Decks zurück.
func (d *Deck) Top() card.Card {
	//Gibt die Oberste Karte aus dem Deck zurück
	return d.Cards[0]
}

// Fügt eine Karte zum Deck hinzu.
func (d *Deck) Add(c card.Card) {

	d.Cards = append([]card.Card{c}, d.Cards...)
}

// Gibt die Anzahl der Karten im Deck zurück.
func (d *Deck) Len() int {
	//Gibt die Länge des Kartenslices im Deck zurück
	return len(d.Cards)
}
