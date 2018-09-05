/*
Have the function CorrectPath(str) read the str parameter being passed, which will represent the movements made in a 5x5 grid of cells starting from the top left position. The characters in the input string will be entirely composed of: r, l, u, d, ?. Each of the characters stand for the direction to take within the grid, for example: r = right, l = left, u = up, d = down. Your goal is to determine what characters the question marks should be in order for a path to be created to go from the top left of the grid all the way to the bottom right without touching previously travelled on cells in the grid.
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

type coordinates struct {
	x int
	y int
}

func solves(str string) bool {
	coords := coordinates{x: 0, y: 0}
	for _, char := range str {
		switch {
		case string(char) == "u":
			coords.y--
		case string(char) == "d":
			coords.y++
		case string(char) == "l":
			coords.x--
		case string(char) == "r":
			coords.x++
		default:
			return false
		}
		// Don't allow negative numbers
		if math.Signbit(float64(coords.x)) || math.Signbit(float64(coords.y)) {
			return false
		}
		// Don't allow numbers higher than 4
		if coords.x > 4 || coords.y > 4 {
			return false
		}
		// TODO don't allow coords which have previously been visited
	}
	if coords.x == 4 && coords.y == 4 {
		return true
	}

	return false
}

func guess(s string, channel chan string) {
	str := []byte(s)
Loop:
	for position, char := range str {
		if string(char) == "?" {
			for _, letter := range []byte{'u', 'd', 'l', 'r'} {
				str[position] = letter
				if strings.Contains(string(str), "?") {
					guess(string(str), channel)
				} else {
					if solves(string(str)) {
						channel <- string(str)
						break Loop
					}
				}
			}
		}
	}
}

func CorrectPath(s string) (solution string) {
	channel := make(chan string)
	go guess(s, channel)
	solution = <-channel

	return
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(CorrectPath(readline()))
	fmt.Println(CorrectPath("r?d?drdd"))

}
