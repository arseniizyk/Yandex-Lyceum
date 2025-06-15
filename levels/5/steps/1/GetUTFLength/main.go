package main

import (
	"testing"
	"unicode/utf8"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		expected int
		wantErr  bool
	}{
		{"with chinese", "Hello, 世界", utf8.RuneCountInString("Hello, 世界"), false},
		{"with smiles", "😁😁😁😁😁😁😁😁", utf8.RuneCountInString("😁😁😁😁😁😁😁😁"), false},
		{"with russian", "Привет, мир", utf8.RuneCountInString("Привет, мир"), false},
		{"invalid", "\xc3\x28", 0, true},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			got, err := GetUTFLength([]byte(cs.text))
			if (err != nil) && !cs.wantErr {
				t.Fatalf("GetUTFLength(%s) = %q, unexpected error: %v", cs.text, got, err)
			}
			if !cs.wantErr && got != cs.expected {
				t.Fatalf("GetUTFLength(%s) = %q, want %q", cs.text, got, cs.expected)
			}
		})
	}
}
