package asciiart

import "fmt"

// PrintWord prints each word in ASCII or a blank line if empty
func PrintWord(words []string, fileLines []string) {
	for _, word := range words {
		// Print blank line for empty word
		if word == "" {
			fmt.Println()
			continue
		}
		// Print ASCII for non-empty word
		PrintASCII(word, fileLines)
	}
}