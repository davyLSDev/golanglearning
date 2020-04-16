package main

// Sum takes an array and returns the sum of all the elements
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}