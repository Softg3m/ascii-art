package main

import (
	"asciiart"
	"fmt"
	"os"
)

func main() {
	// Check if exactly one argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [STRING]")
		return
	}

	input := os.Args[1] // Get input string from command line

	// Exit if input is empty
	if input == "" {
		return
	}

	// Print a single newline if input is literal "\n"
	if input == "\\n" {
		fmt.Println()
		return
	}

	// Read the ASCII font from file
	fileLines, err := asciiart.ReadFont("standard.txt")
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Split input into words/lines
	words := asciiart.ParseInput(input)

	// Print ASCII art for each word/line
	asciiart.PrintWord(words, fileLines)
}