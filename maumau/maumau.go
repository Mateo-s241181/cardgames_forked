package maumau

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
