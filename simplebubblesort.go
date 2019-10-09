// bubblesort has a worst case scenario of needing n-1 iterations

package main

import (
	"fmt"
	"math/rand" // to generate random numbers
	"time"      // to generate a random seed
)

//func sort

func main() {
	var max, min, numberOfNumbers int
	numberOfNumbers = 10
	numsToSortSlice := make([]int, numberOfNumbers)
	const maxItterations = 10 - 1 // array size minus 1
	max = 1000
	min = 1
	rand.Seed(time.Now().UTC().UnixNano())
	for index := 0; index < numberOfNumbers; index++ {
		numsToSortSlice[index] = (rand.Intn(max-min) + min)
	}
	fmt.Println("Slice before sorting is \n", numsToSortSlice)
}

/* Ref: https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
another to help with generating an integer between a min and max value: https://golangcode.com/generate-random-numbers/
*/
