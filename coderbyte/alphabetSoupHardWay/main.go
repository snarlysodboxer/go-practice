/*
implment alphabetical sort manually instead of using the stdlib
*/
package main

import (
	"fmt"
)

var letterValues = map[int]string{
	0:  "a",
	1:  "b",
	2:  "c",
	3:  "d",
	4:  "e",
	5:  "f",
	6:  "g",
	7:  "h",
	8:  "i",
	9:  "j",
	10: "k",
	11: "l",
	12: "m",
	13: "n",
	14: "o",
	15: "p",
	16: "q",
	17: "r",
	18: "s",
	19: "t",
	20: "u",
	21: "v",
	22: "w",
	23: "x",
	24: "y",
	25: "z",
}

func manualSort(str string) string {
	sorted := ""
	for i := 0; i < 26; i++ {
		for _, character := range str {
			if string(character) == letterValues[i] {
				sorted = sorted + letterValues[i]
			}
		}
	}

	return sorted
}

func AlphabetSoup(str string) string {
	return manualSort(str)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(AlphabetSoup("coderbyte"))
	// fmt.Println(AlphabetSoup(readline()))

}
