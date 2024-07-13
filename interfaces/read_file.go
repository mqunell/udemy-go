package main

import (
	"fmt"
	"os"
)

// Read a file and write its contents to the terminal
func read_file_main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Simple solution
	// io.Copy(os.Stdout, file)

	// Explicit solution, in case we need to do more with the file content
	stats, statsErr := file.Stat()
	if statsErr != nil {
		fmt.Println("Error getting file stats:", statsErr)
		os.Exit(1)
	}
	bytes := make([]byte, stats.Size())

	_, readErr := file.Read(bytes)
	if readErr != nil {
		fmt.Println("Error reading file:", readErr)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
