package main

import (
	"reflect"
	"testing"
)

func TestClosestEnemyII(t *testing.T) {
	strArr := []string{"0000", "1000", "0002", "0002"}
	expect := 2
	result := ClosestEnemyII(strArr)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	// contains zero "2"s
	strArr = []string{"0000", "1000", "0000", "0000"}
	expect = 0
	result = ClosestEnemyII(strArr)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	strArr = []string{"000", "100", "200"}
	expect = 1
	result = ClosestEnemyII(strArr)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	strArr = []string{"0000", "2010", "0000", "2002"}
	expect = 2
	result = ClosestEnemyII(strArr)
	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestGetCoordinates(t *testing.T) {
	strArr := []string{"0000", "1000", "0002", "0002"}
	expect := []coordinates{coordinates{1, 0}}
	result := getCoordinates(strArr, "1")
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Expected %v, got %v", expect, result)
	}

	strArr = []string{"0000", "1000", "0002", "0002"}
	expect = []coordinates{coordinates{2, 3}, coordinates{3, 3}}
	result = getCoordinates(strArr, "2")
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}
