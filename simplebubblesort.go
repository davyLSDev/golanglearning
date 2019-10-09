// bubblesort worst case scenario needs n-1 iterations to sort numbers

package main

import (
	"fmt"
	"math/rand" // to generate random numbers
	"time"      // to generate a random seed
)

// fill the slice "numsToSort" with random numbers from 1 to 1000
func randomFill(sliceSize int) []int {
	var max, min int
	max = 1000
	min = 1
	numsToSort := make([]int, sliceSize)
	rand.Seed(time.Now().UTC().UnixNano())
	for index := 0; index < sliceSize; index++ {
		numsToSort[index] = (rand.Intn(max-min) + min)
	}
	return numsToSort
}

func main() {
	var numberOfNumbers int
	numberOfNumbers = 10
	const maxItterations = 10 - 1 // slice size minus 1
	fmt.Println("Slice before sorting is \n", randomFill(numberOfNumbers))
}

/* Ref: https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
another to help with generating an integer between a min and max value: https://golangcode.com/generate-random-numbers/
*/
