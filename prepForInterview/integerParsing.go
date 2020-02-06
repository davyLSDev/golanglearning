package main













// how to scan a line of text from keyboard https://stackoverflow.com/questions/20895552/how-to-read-from-standard-input-in-the-console
// more useful: https://stackoverflow.com/questions/3751429/reading-an-integer-from-standard-input

import (
	"fmt"
	"time"
)

func main() {
	var integer int
	var parsedIntegerNumber []int

	fmt.Println("Enter an integer, and press <RETURN>")
	_, err := fmt.Scanf("%d", &integer)
	fmt.Println("This is your integer ", integer, " with error ", err)

	t0 := time.Now()
	parsedIntegerNumber = integerParsing(integer)
	t1 := time.Now()
	fmt.Println("Time for string converted integer parsing ", t1.Sub(t0))
	fmt.Println("This is the parsed integer ", parsedIntegerNumber)
}

// had to look up about returning a slice, but it makes sense . . .
// how to prepend a slice https://codingair.wordpress.com/2014/07/18/go-appendprepend-item-into-slice/
func integerParsing(number int) (parsedNumber []int) {
	var digit int
	for number > 0 {
		digit = number % 10
		parsedNumber = append([]int{digit}, parsedNumber...)
		number = number/10
	}
	return parsedNumber
}
