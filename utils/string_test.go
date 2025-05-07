package utils

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple ASCII string",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Unicode string",
			input:    "привет",
			expected: "тевирп",
		},
		{
			name:     "String with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
} 