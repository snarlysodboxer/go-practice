package main

import (
	"reflect"
	"testing"
)

func TestSumAndParts(t *testing.T) {
	example := []int{1, 2, 3}
	expected := map[int][]int{
		1: []int{1},
		2: []int{2},
		3: []int{3},
		4: []int{1, 3},
		5: []int{2, 3},
	}
	sumParts := sumAndParts(example)
	if !reflect.DeepEqual(sumParts, expected) {
		t.Errorf("Expected %v, got %v", expected, sumParts)
	}
}

func TestScaleWeights(t *testing.T) {
	example := []string{"[3, 4]", "[1, 2, 7, 7]"}
	expectedScale := []int{3, 4}
	expectedWeights := []int{1, 2, 7, 7}
	scale, weights := scaleWeights(example)
	if !reflect.DeepEqual(expectedScale, scale) || !reflect.DeepEqual(expectedWeights, weights) {
		t.Errorf("Expected scale %v, got %v", expectedScale, scale)
		t.Errorf("Expected weights %v got %v", expectedWeights, weights)
	}
}

func TestScaleBalancing(t *testing.T) {
	example := []string{"[3, 4]", "[1, 2, 7, 7]"}
	if str := ScaleBalancing(example); str != "1" {
		t.Errorf("Expected '1', got %s", str)
	}

	example = []string{"[13, 4]", "[1, 2, 3, 6, 14]"}
	if str := ScaleBalancing(example); str != "3,6" {
		t.Errorf("Expected '3,6', got %s", str)
	}

	example = []string{"[5, 9]", "[1, 2]"}
	if str := ScaleBalancing(example); str != "not found" {
		t.Errorf("Expected 'not found', got %s", str)
	}

	example = []string{"[5, 9]", "[1, 2, 6, 7]"}
	if str := ScaleBalancing(example); str != "2,6" {
		t.Errorf("Expected '2,6', got %s", str)
	}
}
