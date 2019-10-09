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

func simpleSort(sortedNumbers []int) []int {
	var temp int
	maxItterations := len(sortedNumbers) - 1 // worst case scenario
	numbersToSort := maxItterations

	for itterations := 0; itterations < maxItterations; itterations++ {
		for index := 0; index < numbersToSort; index++ {
			if sortedNumbers[index] > sortedNumbers[index+1] {
				temp = sortedNumbers[index]
				sortedNumbers[index] = sortedNumbers[index+1]
				sortedNumbers[index+1] = temp
			}
		}
	}

	return sortedNumbers
}

func main() {
	var numberOfNumbers int
	numberOfNumbers = 10
	numbers := make([]int, numberOfNumbers)
	const maxItterations = 10 - 1 // slice size minus 1
	numbers = randomFill(numberOfNumbers)
	fmt.Println("Slice before sorting is \n", numbers)
	fmt.Println("Slice after sort is \n", simpleSort(numbers))
}

/* Ref: https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
another to help with generating an integer between a min and max value: https://golangcode.com/generate-random-numbers/
*/
