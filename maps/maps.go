package main

import "fmt"

func main() {
	// Multiple ways of declaring and initializing:
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}
	colors["blue"] = "#0000ff" // Always [] syntax with maps
	delete(colors, "green")

	fmt.Println(colors) // map[blue:#0000ff red:#ff0000]
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Printf("%v: %v\n", key, val)
	}
}
