// My understanding of this is that the algorithm has a sorted slice as the input
// and returns the index of the value that is being searched for
// see https://www.codementor.io/learn-programming/3-essential-algorithm-examples-you-should-know

package main

import (
	"fmt"
)

// having trouble recalling syntax of a function!
// ahhh yes, don't put a comma after "input" and before "[]int"
func binarySearch(input []int, integerToFind int) (index int) {
	index = integerToFind
	return
}

func main() {
	var intToFind int
	sorted := []int{10, 12, 33, 54, 65, 96, 107} // I had trouble remembering this syntax!
	fmt.Println(sorted)
	fmt.Println("The index of the number being searched for, is ", binarySearch(sorted, intToFind))
}
