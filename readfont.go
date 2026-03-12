package asciiart

import (
	"os"
	"strings"
)

// ReadFont reads a font file and returns its lines
func ReadFont(filename string) ([]string, error) {
	data, err := os.ReadFile(filename) // Read entire file
	if err != nil {
		return nil, err // Return error if reading fails
	}
	return strings.Split(string(data), "\n"), nil // Split file into lines
}