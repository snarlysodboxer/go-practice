/*
Have the function ClosestEnemyII(strArr) read the matrix of numbers stored in strArr which will be a 2D matrix that contains only the integers 1, 0, or 2. Then from the position in the matrix where a 1 is, return the number of spaces either left, right, down, or up you must move to reach an enemy which is represented by a 2. You are able to wrap around one side of the matrix to the other as well.
*/
package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type coordinates struct {
	line     int
	position int
}

func getCoordinates(strArr []string, str string) []coordinates {
	results := []coordinates{}
	for linePosition, line := range strArr {
		coords := coordinates{}
		if strings.Contains(string(line), str) {
			for charPosition, char := range line {
				if strings.Contains(string(char), str) {
					coords.line = linePosition
					coords.position = charPosition
					results = append(results, coords)
				}
			}
		}
	}

	return results
}

func ClosestEnemyII(strArr []string) int {
	// ensure strArr contains at least one "2"
	all := strings.Join(strArr, "")
	if !strings.Contains(all, "2") {
		return 0
	}

	// find the coordinates for the 1
	oneCoords := getCoordinates(strArr, "1")
	if len(oneCoords) != 1 {
		panic("More than one 1!")
	}

	// find all 2 coordinates
	twoCoords := getCoordinates(strArr, "2")

	// find position differences
	answers := []int{}
	for _, coords := range twoCoords {
		upDownDiff := int(math.Abs(float64(coords.line - oneCoords[0].line)))

		noWrap := int(math.Abs(float64(oneCoords[0].position - coords.position)))
		wrap := oneCoords[0].position + int((math.Abs(float64(coords.position - len(strArr[0])))))
		leftRightDiff := 0
		if noWrap < wrap {
			leftRightDiff = noWrap
		} else {
			leftRightDiff = wrap
		}

		answers = append(answers, upDownDiff+leftRightDiff)
	}

	// find lowest answer
	sort.Ints(answers)

	return answers[0]

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(ClosestEnemyII(readline()))
	fmt.Println(ClosestEnemyII([]string{"0000", "1000", "0002", "0002"}))

}
