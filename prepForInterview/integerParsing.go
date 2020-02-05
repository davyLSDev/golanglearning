package main





// how to scan a line of text from keyboard https://stackoverflow.com/questions/20895552/how-to-read-from-standard-input-in-the-console
// more useful: https://stackoverflow.com/questions/3751429/reading-an-integer-from-standard-input

import (
	"fmt"
	"time"
)

func main() {
	var integer int
	t0 := time.Now()
	fmt.Println("Enter an integer, and press <RETURN>")
	_, err := fmt.Scanf("%d", &integer)
	fmt.Println("This is your integer ", integer, " with error ", err)

	t1 := time.Now()
	fmt.Println("Time for string converted integer parsing ", t1.Sub(t0))
}
