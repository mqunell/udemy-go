package main

var deckSize int = 52

func main() {
	deck := newDeckFromFile("deck.txt")
	hand, deck := deal(deck, 5)

	hand.print()
	deck.print()
}

// $ go run main.go deck.go
