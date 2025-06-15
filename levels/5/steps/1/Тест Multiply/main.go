package main

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "with positive numbers",
			a:        5,
			b:        5,
			expected: 25,
		},
		{
			name:     "with negative numbers",
			a:        -3,
			b:        -5,
			expected: 15,
		},
		{
			name:     "with mixed numbers",
			a:        -5,
			b:        5,
			expected: -25,
		},
		{
			name:     "with zero",
			a:        0,
			b:        5,
			expected: 0,
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			got := Multiply(cs.a, cs.b)
			if got != cs.expected {
				t.Fatalf("Length(%d, %d) = %q, want %q", cs.a, cs.b, got, cs.expected)
			}
		})
	}

}
