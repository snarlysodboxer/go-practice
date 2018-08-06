package main

import (
	"testing"
)

func TestIsVowel(t *testing.T) {
	chars := []string{"a", "e", "i", "o", "u", "y"}
	for _, char := range chars {
		if truthy := isVowel(char); !truthy {
			t.Errorf("Expected true, got %t", truthy)
		}
	}

	chars = []string{"b", "c", "g", "r", "w", "z"}
	for _, char := range chars {
		if truthy := isVowel(char); truthy {
			t.Errorf("Expected false, got %t", truthy)
		}
	}
}

func TestVowelSquare(t *testing.T) {
	// more than 1 square
	matrix := []string{
		"aqrst",
		"ukaei",
		"ffooo",
	}
	expect := "1-2"
	square := VowelSquare(matrix)
	if square != expect {
		t.Errorf("Expected '%s', got '%s'", expect, square)
	}

	// position 0
	matrix = []string{
		"abcd",
		"eikr",
		"oufj",
	}
	expect = "1-0"
	square = VowelSquare(matrix)
	if square != expect {
		t.Errorf("Expected '%s', got '%s'", expect, square)
	}

	// not found
	matrix = []string{
		"gg",
		"ff",
	}
	expect = "not found"
	square = VowelSquare(matrix)
	if square != expect {
		t.Errorf("Expected '%s', got '%s'", expect, square)
	}
}
