package main

var deckSize int = 52

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// $ go run main.go deck.go
