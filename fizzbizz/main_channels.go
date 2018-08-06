/*
"Write a program that prints the numbers from 1 to 100. But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”. For numbers which are multiples of both three and five print “FizzBuzz”."
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)

	go func(ch chan string) {
		for line := range ch {
			fmt.Println(line)
		}
	}(ch)

	for i := 1; i <= 100; i++ {
		d3 := i%3 == 0
		d5 := i%5 == 0

		switch {
		case d3 && d5:
			ch <- "FizzBuzz"
		case d3:
			ch <- "Fizz"
		case d5:
			ch <- "Buzz"
		default:
			ch <- strconv.Itoa(i)
		}

	}

	close(ch)
}
