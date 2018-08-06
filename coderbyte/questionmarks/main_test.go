package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestFindSets(t *testing.T) {
	// only 1 set
	str := "1a6?9"
	expect := [][]string{[]string{"1", "a", "6", "?", "9"}}
	sets := findSets(str)
	if !reflect.DeepEqual(sets, expect) {
		t.Errorf("Expected %v, got %v", expect, sets)
	}

	// multiple sets non-overlapping
	str = "1a6?9fds5fds5"
	expect = [][]string{[]string{"1", "a", "6", "?", "9"}, []string{"5", "f", "d", "s", "5"}}
	sets = findSets(str)
	if !reflect.DeepEqual(sets, expect) {
		t.Errorf("Expected %v, got %v", expect, sets)
	}

	// multiple sets overlapping
	str = "acc?7??sss?3r?3r5???5ASDF"
	expect = [][]string{[]string{"7", "?", "?", "s", "s", "s", "?", "3"}, []string{"7", "?", "?", "s", "s", "s", "?", "3", "r", "?", "3"}, []string{"5", "?", "?", "?", "5"}}
	sets = findSets(str)
	if !reflect.DeepEqual(sets, expect) {
		t.Errorf("Expected %v, got %v", expect, sets)
	}
}

func TestThreeQMarks(t *testing.T) {
	// only 1 question mark
	set := strings.Split("1a6?9", "")
	if truthy := threeQMarks(set); truthy != false {
		t.Errorf("Expected false, got %t", truthy)
	}

	// 3 question marks
	set = strings.Split("7cc?7??sss3", "")
	if truthy := threeQMarks(set); truthy != true {
		t.Errorf("Expected true, got %t", truthy)
	}

	// 4 question marks
	set = strings.Split("4acc?7??s?6", "")
	if truthy := threeQMarks(set); truthy != false {
		t.Errorf("Expected false, got %t", truthy)
	}
}

func TestQuestionsMarks(t *testing.T) {
	str := "aa6?9"
	if truthy := QuestionsMarks(str); truthy != "false" {
		t.Errorf("Expected 'false', got %s", truthy)
	}

	str = "acc?7??sss?3rr1??????5"
	if truthy := QuestionsMarks(str); truthy != "true" {
		t.Errorf("Expected 'true', got %s", truthy)
	}

	str = "acc?7??sss?3rr5???5ASDF"
	if truthy := QuestionsMarks(str); truthy != "true" {
		t.Errorf("Expected 'true', got %s", truthy)
	}

	str = "acc?7??sss?3r?3r5???5ASDF"
	if truthy := QuestionsMarks(str); truthy != "false" {
		t.Errorf("Expected 'false', got %s", truthy)
	}
}
