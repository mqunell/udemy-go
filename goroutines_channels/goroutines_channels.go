package main

import (
	"fmt"
	"net/http"
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
	for {
		go checkLink(c, <-c)
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
