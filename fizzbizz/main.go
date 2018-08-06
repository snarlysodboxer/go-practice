/*
"Write a program that prints the numbers from 1 to 100. But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”. For numbers which are multiples of both three and five print “FizzBuzz”."
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var toPrint []string
	for i := 1; i <= 100; i++ {
		d3 := i%3 == 0
		d5 := i%5 == 0

		switch {
		case d3 && d5:
			toPrint = append(toPrint, "FizzBuzz")
		case d3:
			toPrint = append(toPrint, "Fizz")
		case d5:
			toPrint = append(toPrint, "Buzz")
		default:
			toPrint = append(toPrint, strconv.Itoa(i))
		}

	}
	for _, line := range toPrint {
		fmt.Println(line)
	}
}
