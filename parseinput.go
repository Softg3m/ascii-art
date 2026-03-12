package asciiart

import "strings"

// ParseInput splits input into lines and trims leading spaces
func ParseInput(input string) []string {
	// Split on literal "\n"
	words := strings.Split(input, `\n`)

	for i, word := range words {
		// Remove leading spaces from the word and align ASCII output to the left
		words[i] = strings.TrimLeft(word, " ")
	}

	return words
}