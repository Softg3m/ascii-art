package asciiart

import (
	"testing"
	"reflect"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single word",
			input:    "Hello",
			expected: []string{"Hello"},
		},
		{
			name:     "two words split by newline",
			input:    `Hello\nWorld`,
			expected: []string{"Hello", "World"},
		},
		{
			name:     "multiple newlines",
			input:    `Hello\nWorld\nFoo`,
			expected: []string{"Hello", "World", "Foo"},
		},
		{
			name:     "leading spaces trimmed after newline",
			input:    `Hello\n   World`,
			expected: []string{"Hello", "World"},
		},
		{
			name:     "leading spaces trimmed on first word",
			input:    "   Hello",
			expected: []string{"Hello"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "only newline sequence",
			input:    `\n`,
			expected: []string{"", ""},
		},
		{
			name:     "multiple consecutive newlines",
			input:    `Hello\n\nWorld`,
			expected: []string{"Hello", "", "World"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseInput(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseInput(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}