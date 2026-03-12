package asciiart

import "fmt"

// PrintASCII prints a single word in ASCII, row by row
func PrintASCII(word string, fileLines []string) {
	for j := 1; j <= 8; j++ { // Loop over each of the 8 rows of ASCII
		for _, char := range word { // Loop over each character
			asciiIndex := int(char) - 32   // Map character to font index
			start := asciiIndex * 9        // Calculate start line in font file
			fmt.Print(fileLines[start+j])  // Print the correct row of this character
		}
		fmt.Println() // Newline after each row
	}
}