/*
Have the function VowelSquare(strArr) take the strArr parameter being passed which will be a 2D matrix of some arbitrary size filled with letters from the alphabet, and determine if a 2x2 square composed entirely of vowels exists in the matrix.
*/
package main

import "fmt"

func isVowel(str string) bool {
	chars := []string{"a", "e", "i", "o", "u", "y"}
	for _, char := range chars {
		if str == char {
			return true
		}
	}
	return false
}

func VowelSquare(strArr []string) string {
	result := "not found"
	for arrayPosition, str := range strArr {
		if arrayPosition == len(strArr)-1 {
			continue
		}
		for charPosition, char := range str {
			if charPosition == len(str)-1 {
				continue
			}
			if isVowel(string(char)) && isVowel(string(str[charPosition+1])) {
				if isVowel(string(strArr[arrayPosition+1][charPosition])) && isVowel(string(strArr[arrayPosition+1][charPosition+1])) {
					result = fmt.Sprintf("%d-%d", arrayPosition, charPosition)
					break
				}
			}
		}
	}

	return result

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(VowelSquare(readline()))
	fmt.Println(VowelSquare([]string{"aqrst", "ukaei", "ffooo"}))

}
