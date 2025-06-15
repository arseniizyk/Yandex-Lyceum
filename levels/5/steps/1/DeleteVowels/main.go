package main

import (
	"testing"
)

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "with o and O",
			text:     "Woooooo! WO",
			expected: "W! W",
		},
		{
			name:     "with a and A",
			text:     "YAaaaaaa, Paaaa",
			expected: "Y, P",
		},
		{
			name:     "with i and I",
			text:     "RiIi, Liii",
			expected: "R, L",
		},
		{
			name:     "with e and E",
			text:     "Ser, Her, ShE",
			expected: "Sr, Hr, Sh",
		},
		{
			name:     "with u and U",
			text:     "TUUuuuu, Twu",
			expected: "T, Tw",
		},
		{
			name:     "mixed",
			text:     "Lorem ipsum testing this",
			expected: "Lrm psm tstng ths",
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			got := DeleteVowels(cs.text)
			if got != cs.expected {
				t.Fatalf("DeleteVowels(%s) = %q, want %q", cs.text, got, cs.expected)
			}
		})
	}
}
