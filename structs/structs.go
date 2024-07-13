package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	name string
	age  int
	contactInfo
}

func main() {
	// Multiple ways of declaring and initializing:
	// var matt person
	// matt := person{"Matt", 30, contactInfo{"matt@github.com", 98000}}

	matt := person{
		name: "Matt",
		age:  30,
		contactInfo: contactInfo{
			email: "matt@github.com",
			zip:   98000,
		},
	}

	fmt.Printf("%+v\n", matt) // {name:Matt age:30 contactInfo:{email:matt@github.com zip:98000}}
}
