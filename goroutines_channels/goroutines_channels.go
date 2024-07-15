package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",
		"http://google.com",
		"http://stackoverflow.com",
	}

	// Create a channel that can transfer strings
	c := make(chan string)

	for _, link := range links {
		// Run checkLink within a goroutine
		go checkLink(c, link)
	}

	// Wait for the results (one iteration for each link)
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Call checkLink again every time the channel receives a message
	// for {
	// 	go checkLink(c, <-c)
	// }

	// Effectively the same loop definition as above, but more clear
	// Values received by `c` are assigned to `link`
	for link := range c {
		// It's important to pass `link` as an arg to the function literal so it's copied
		// (because Go is pass-by-value) and the goroutine doesn't reference a `link` var
		// that is shared with another goroutine
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(c, l)
		}(link)
	}
}

func checkLink(c chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "appears to be down")
		c <- link
		return
	}

	fmt.Println(link, "is working")
	c <- link
}
