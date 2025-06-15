package main

import (
	"testing"
)

func TestLength(t *testing.T) {
	cases := []struct {
		name     string
		number   int
		expected string
	}{
		{
			name:     "negative",
			number:   -1,
			expected: "negative",
		},
		{
			name:     "zero",
			number:   0,
			expected: "zero",
		},
		{
			name:     "short",
			number:   9,
			expected: "short",
		},
		{
			name:     "long",
			number:   99,
			expected: "long",
		},
		{
			name:     "very long",
			number:   10000,
			expected: "very long",
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			got := Length(cs.number)
			if got != cs.expected {
				t.Fatalf("Length(%d) = %q, want %q", cs.number, got, cs.expected)
			}
		})
	}
}
