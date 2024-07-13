package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	// Pass by value types (int, float, string, bool, struct)
	// need pointers to be changed within another function
	matt := person{
		name: "Matt",
		age:  30,
	}

	matt.print()            // 1. {name:Matt age:30}
	matt.updateName("Thew") // Shortcut for (&matt).updateName("Thew")
	matt.print()            // 3. {name:Thew age:30}

	// ---

	// Pass by reference types (slice, map, channel, pointer, function)
	// do not need pointers to be changed within another function
	s := []string{"A", "B", "C"}
	fmt.Println(s) // [A B C]
	updateSlice(s)
	fmt.Println(s) // [X Y Z]
}

// (p *person) means pointer to a person
func (p *person) updateName(newName string) {
	p.name = newName // Shortcut for (*p).name = newName
	p.print()        // 2. {name:Thew age:30} - also shortcut for (*p).print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateSlice(s []string) {
	s[0] = "X"
	s[1] = "Y"
	s[2] = "Z"
}
