package main

import "fmt"

// By using this interface type as a parameter, a function can guarantee that
// its received argument will be some type that has these methods. englishBot
// and spanishBot don't have to explicitly "implement" the interface with any
// keywords; Go detects that their respective `getGreeting` functions satisfy
// this contract and considers them to be `bot`s automatically.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// Imagine there is logic here specific for englishBot
	return "hello"
}

func (spanishBot) getGreeting() string {
	// Imagine there is logic here specific for spanishBot
	return "hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
