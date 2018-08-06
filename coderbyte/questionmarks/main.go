/*
From https://www.coderbyte.com/editor/guest:Questions%20Marks:Go

Have the function QuestionsMarks(str) take the str string parameter, which will contain single digit numbers, letters, and question marks, and check if there are exactly 3 question marks between every pair of two numbers that add up to 10. If so, then your program should return the string true, otherwise it should return the string false. If there aren't any two numbers that add up to 10 in the string, then your program should return false as well.
*/
package main

import (
	"fmt"
	"strconv"
)

func findNext(startPosition int, str string) (position, number int, isMore bool) {
	var char rune
	var err error
	isMore = false
	for position, char = range str {
		if position <= startPosition {
			continue
		}
		number, err = strconv.Atoi(string(char))
		if err == nil {
			break
		}
	}
	oldPosition := position
	for pos, c := range str {
		if pos <= oldPosition {
			continue
		}
		_, err = strconv.Atoi(string(c))
		if err == nil {
			isMore = true
		}
	}

	return
}

func exactly3QMarks(oldPosition, position int, str string) bool {
	var qMarks []string
	for ; oldPosition < position; oldPosition++ {
		if string(str[oldPosition]) == "?" {
			qMarks = append(qMarks, string(str[oldPosition]))
		}

	}

	if len(qMarks) == 3 {
		return true
	}

	return false
}

func QuestionsMarks(str string) string {
	truthy := "false"
	position := 0
	number := 0
	isMore := true
	for isMore {
		lastNumber := number
		lastPosition := position
		position, number, isMore = findNext(position, str)
		if lastNumber+number == 10 {
			if exactly3QMarks(lastPosition, position, str) {
				truthy = "true"
			} else {
				truthy = "false"
			}
		}
	}

	return truthy
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(QuestionsMarks(readline()))
	fmt.Println(QuestionsMarks("acc?7??sss?3rr1??????5"))

}
