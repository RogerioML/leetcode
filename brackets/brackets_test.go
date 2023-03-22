package brackets

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "valid string with one bracket",
			input:    "()",
			expected: true,
		},
		{
			name:     "valid string with multiple brackets",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "valid string with nested brackets",
			input:    "{[]}",
			expected: true,
		},
		{
			name:     "invalid string with unmatching brackets",
			input:    "(]",
			expected: false,
		},
		{
			name:     "invalid string with incomplete brackets",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "invalid string with characters other than brackets",
			input:    "{a[b]c}d",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsValid(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v %v", tc.expected, result, tc.input)
			}
		})
	}
}
