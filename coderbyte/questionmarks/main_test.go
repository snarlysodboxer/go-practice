package main

import (
	"testing"
)

func TestFindNext(t *testing.T) {
	str := "acc?7??sss?3rr1??????5"
	position := 0
	number := 0
	isMore := true
	if position, number, isMore = findNext(position, str); position != 4 || number != 7 || isMore != true {
		t.Errorf("Expected position 4, got %d", position)
		t.Errorf("Expected number 7, got %d", number)
		t.Errorf("Expected isMore to be true, got %t", isMore)
	}

	if position, number, isMore = findNext(position, str); position != 11 || number != 3 || isMore != true {
		t.Errorf("Expected position 11, got %d", position)
		t.Errorf("Expected number 3, got %d", number)
		t.Errorf("Expected isMore to be true, got %t", isMore)
	}

	if position, number, isMore = findNext(position, str); position != 14 || number != 1 || isMore != true {
		t.Errorf("Expected position 14, got %d", position)
		t.Errorf("Expected number 1, got %d", number)
		t.Errorf("Expected isMore to be true, got %t", isMore)
	}

	if position, number, isMore = findNext(position, str); position != 21 || number != 5 || isMore != false {
		t.Errorf("Expected position 21, got %d", position)
		t.Errorf("Expected number 5, got %d", number)
		t.Errorf("Expected isMore to be false, got %t", isMore)
	}

	str = "acc?7??sss?"
	position = 0
	number = 0
	isMore = true
	if position, number, isMore = findNext(position, str); position != 4 || number != 7 || isMore != false {
		t.Errorf("Expected position 4, got %d", position)
		t.Errorf("Expected number 7, got %d", number)
		t.Errorf("Expected isMore to be false, got %t", isMore)
	}
}

func TestExactly3QMarks(t *testing.T) {
	// only 1 question mark
	str := "aa6?9"
	if truthy := exactly3QMarks(2, 4, str); truthy != false {
		t.Errorf("Expected false, got %t", truthy)
	}

	// 3 question marks
	str = "acc?7??sss?3rr1??????5"
	if truthy := exactly3QMarks(4, 11, str); truthy != true {
		t.Errorf("Expected true, got %t", truthy)
	}

	// 4 question marks
	str = "acc?7??s?ss?3rr1??????5"
	if truthy := exactly3QMarks(5, 12, str); truthy != false {
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
}
