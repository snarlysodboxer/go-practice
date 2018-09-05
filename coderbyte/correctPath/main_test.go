package main

import (
	"testing"
)

func TestGuess(t *testing.T) {
	input := "r?d?drdd"
	expect := "rrdrdrdd"
	result := CorrectPath(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	input = "???rrurdr?"
	expect = "dddrrurdrd"
	result = CorrectPath(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	input = "drdr??rrddd?"
	expect = "drdruurrdddd"
	result = CorrectPath(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestSolves(t *testing.T) {
	input := "rrdrdrdd"
	expect := true
	result := solves(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	input = "urdrdrdd"
	expect = false
	result = solves(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	input = "drdrrrrrdddr"
	expect = false
	result = solves(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	input = "rrrrrurdrr"
	expect = false
	result = solves(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	input = "dddrrurdrd"
	expect = true
	result = solves(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	input = "drdruurrdddd"
	expect = true
	result = solves(input)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}
