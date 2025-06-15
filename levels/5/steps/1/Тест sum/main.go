package main

import "testing"

func TestSum(t *testing.T) {
	got := Sum(1, 5)
	expected := 6
	if got != expected {
		t.Fatalf("Sum(1,5) = %q, want %q", got, expected)
	}
}
