package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Implicitly implements the io.Writer interface
type logWriter struct{}

func http_main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Example 1: Making a byte array of arbitrary length for Read
	bytes := make([]byte, 99999)
	resp.Body.Read(bytes)
	fmt.Println(string(bytes))

	// Example 2: Printing via `os.Stdout`
	io.Copy(os.Stdout, resp.Body)

	// Example 3: Printing via custom io.Writer interface
	l, err := io.Copy(logWriter{}, resp.Body)
	fmt.Println(l, err)
}

func (logWriter) Write(bytes []byte) (int, error) {
	fmt.Println(string(bytes))
	return len(bytes), nil
}
