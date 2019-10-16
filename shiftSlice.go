// shift slice is a software to shift the contents of an aray right and left
package main

import (
	"fmt"
	"math/rand"
	"time"
	/*	"math/rand"
		"time"*/)

// fill the slice "intSlice" with random numbers from 1 to 1000
func randomFill(size int) []int {
	var max, min int
	max = 1000
	min = 1
	intSlice := make([]int, size)
	rand.Seed(time.Now().UTC().UnixNano())
	for index := 0; index < size; index++ {
		intSlice[index] = (rand.Intn(max-min) + min)
	}
	return intSlice
}

func main() {
	sliceSize := 10
	/*	shifts := 2
		left := false
		right := true */
	integersSlice := make([]int, sliceSize)
	integersSlice = randomFill(sliceSize)
	fmt.Println(integersSlice)
	//	shiftSlice(integersSlice, shifts, left)
}
