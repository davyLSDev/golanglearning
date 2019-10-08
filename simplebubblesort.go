// bubblesort has a worst case scenario of needing n-1 iterations

package main

import (
	"fmt"
	"math/rand"	// to generate random numbers
	"time" 			// to generate a random seed
)

func main () {
	var max, min int
	max = 1000
	min =1
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println ("The random number is", (rand.Intn(max - min) + min))
}

/* Ref: https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
another to help with generating an integer between a min and max value: https://golangcode.com/generate-random-numbers/
*/