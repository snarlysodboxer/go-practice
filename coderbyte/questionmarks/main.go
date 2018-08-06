/*
From https://www.coderbyte.com/editor/guest:Questions%20Marks:Go

Have the function QuestionsMarks(str) take the str string parameter, which will contain single digit numbers, letters, and question marks, and check if there are exactly 3 question marks between every pair of two numbers that add up to 10. If so, then your program should return the string true, otherwise it should return the string false. If there aren't any two numbers that add up to 10 in the string, then your program should return false as well.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func threeQMarks(chars []string) bool {
	count := 0
	for _, char := range chars {
		if char == "?" {
			count++
		}

	}

	if count == 3 {
		return true
	}

	return false
}

func findSets(str string) (sets [][]string) {
	for position, char := range str {
		num, err := strconv.Atoi(string(char))
		if err == nil {
			for nextPosition, nextChar := range str[position+1:] {
				nextNum, err := strconv.Atoi(string(nextChar))
				if err == nil {
					if num+nextNum == 10 {
						set := strings.Split(str[position:position+nextPosition+2], "")
						sets = append(sets, set)
					}
				}
			}
		}
	}

	return
}

func QuestionsMarks(str string) string {
	truthy := "false"

	sets := findSets(str)
	for _, set := range sets {
		if threeQMarks(set) {
			truthy = "true"
		} else {
			return "false"
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
