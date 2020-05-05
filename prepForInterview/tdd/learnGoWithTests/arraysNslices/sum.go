package main

// Sum takes an array and returns the sum of all the elements
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

// SumAll takes integer arrays and returns the sums of all the elements of each array in an array
func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers) // Note, there is an error in the tutorial "sums := ..."

	for idx, numbers := range numbersToSum {
		sums[idx] = Sum(numbers)
	}
	return
}

// SumAllTails takes the integer arrays and returns the sums of all the elements except the "head" of each array and puts these sums into an array */
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
