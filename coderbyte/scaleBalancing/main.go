/*
From https://www.coderbyte.com/editor/guest:Scale%20Balancing:Go

Have the function ScaleBalancing(strArr) read strArr which will contain two elements, the first being the two positive integer weights on a balance scale (left and right sides) and the second element being a list of available weights as positive integers. Your goal is to determine if you can balance the scale by using the least amount of weights from the list, but using at most only 2 weights. For example: if strArr is ["[5, 9]", "[1, 2, 6, 7]"] then this means there is a balance scale with a weight of 5 on the left side and 9 on the right side. It is in fact possible to balance this scale by adding a 6 to the left side from the list of weights and adding a 2 to the right side. Both scales will now equal 11 and they are perfectly balanced. Your program should return a comma separated string of the weights that were used from the list in ascending order, so for this example your program should return the string 2,6

There will only ever be one unique solution and the list of available weights will not be empty. It is also possible to add two weights to only one side of the scale to balance it. If it is not possible to balance the scale then your program should return the string not possible.

*/
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func scaleWeights(strArr []string) (scale, weights []int) {
	fieldsFunc := func(char rune) bool {
		return !unicode.IsNumber(char)
	}

	scaleStrSlice := strings.FieldsFunc(strArr[0], fieldsFunc)
	for _, char := range scaleStrSlice {
		number, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		scale = append(scale, number)
	}

	weightsStrSlice := strings.FieldsFunc(strArr[1], fieldsFunc)
	for _, char := range weightsStrSlice {
		number, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		weights = append(weights, number)
	}

	return scale, weights
}

func sumAndParts(weights []int) map[int][]int {
	results := map[int][]int{}
	for _, weight := range weights {
		results[weight] = []int{weight}
	}
	for position, weight := range weights {
		for _, w := range weights[position+1:] {
			sum := w + weight
			if _, ok := results[sum]; !ok {
				results[sum] = []int{weight, w}
			}
		}
	}

	return results
}

func ScaleBalancing(strArr []string) string {
	// turn strArr into 2 int slices
	scale, weights := scaleWeights(strArr)

	// calculate all sums creating a map of the sum to it's parts
	sumParts := sumAndParts(weights)

	// find lowest number/side
	lowPosition := 3
	highPosition := 3
	if scale[0] > scale[1] {
		lowPosition = 1
		highPosition = 0
	} else {
		lowPosition = 0
		highPosition = 1
	}

	// loop through sumParts, find one sided balancers
	result := "not found"
	for sum, parts := range sumParts {
		if sum+scale[lowPosition] == scale[highPosition] {
			result = strings.Trim(strings.Replace(fmt.Sprint(parts), " ", ",", -1), "[]")
		}
	}

	if result == "not found" {
		// find balances that include adding one weight to one side and another
		//   weight to the other side
		// for each weight, add it to one side, and one by one add the remaining weights to the other side, and see if sums match
		for position, weight := range weights {
			for p, w := range weights {
				if p == position {
					continue
				}
				if scale[0]+weight == scale[1]+w || scale[0]+w == scale[1]+weight {
					result = fmt.Sprintf("%d,%d", w, weight)
				}

			}
		}
	}

	return result
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(ScaleBalancing(readline()))
	fmt.Println(ScaleBalancing([]string{"[3, 4]", "[1, 2, 7, 7]"}))
}
