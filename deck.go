package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three"}

	cards := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// TODO: Try rewriting this as a method with `(d *deck)` as the receiver after covering pointers
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(fileName string) error {
	bytes := []byte(d.toString())
	err := os.WriteFile(fileName, bytes, 0666)
	return err
}

func newDeckFromFile(fileName string) deck {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	cards := strings.Split(string(bytes), "\n")
	return deck(cards)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
