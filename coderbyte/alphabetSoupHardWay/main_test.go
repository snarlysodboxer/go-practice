package main

import "testing"

func TestManualSort(t *testing.T) {
	sorted := manualSort("coderbyte")
	expected := "bcdeeorty"
	if sorted != expected {
		t.Fatalf("Expected %s, got %s", expected, sorted)
	}

	sorted = manualSort("aacoderbyte")
	expected = "aabcdeeorty"
	if sorted != expected {
		t.Fatalf("Expected %s, got %s", expected, sorted)
	}
}
