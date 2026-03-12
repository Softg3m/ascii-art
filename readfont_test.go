// readfont_test.go
package asciiart

import (
	"os"
	"testing"
)

func TestReadFont(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expectLen   int
		expectFirst string
		expectError bool
	}{
		{
			name:        "valid file with newlines",
			content:     "line1\nline2\nline3",
			expectLen:   3,
			expectFirst: "line1",
			expectError: false,
		},
		{
			name:        "single line file",
			content:     "onlyline",
			expectLen:   1,
			expectFirst: "onlyline",
			expectError: false,
		},
		{
			name:        "empty file",
			content:     "",
			expectLen:   1,
			expectFirst: "",
			expectError: false,
		},
		{
			name:        "file not found",
			content:     "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectError {
				_, err := ReadFont("nonexistent_file.txt")
				if err == nil {
					t.Error("expected error for missing file, got nil")
				}
				return
			}

			// Write a temp file
			tmp, err := os.CreateTemp("", "font_test_*.txt")
			if err != nil {
				t.Fatal("failed to create temp file:", err)
			}
			defer os.Remove(tmp.Name())
			tmp.WriteString(tt.content)
			tmp.Close()

			lines, err := ReadFont(tmp.Name())
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(lines) != tt.expectLen {
				t.Errorf("got %d lines, want %d", len(lines), tt.expectLen)
			}
			if lines[0] != tt.expectFirst {
				t.Errorf("first line = %q, want %q", lines[0], tt.expectFirst)
			}
		})
	}
}