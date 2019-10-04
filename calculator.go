package main

import (
// a great way to remember to import a package before using it       _ "fmt"
	"fmt"
	"os"
)

// names for operands here: https://math.stackexchange.com/questions/975541/what-are-the-formal-names-of-operands-and-results-for-basic-operations

func main() {
	argCount := len(os.Args[1:])
	operation := os.Args[2]
	fmt.Println("There are ", argCount, "arguments entered")
	fmt.Println("This is a very simplisitic way to do command parsing")
	fmt.Println("You want to do the following arithmetic operation: ", operation)
	fmt.Println("For integer arithmetic")
	fmt.Println("Adding function 1 + 2 is: ", add(1,2))
	fmt.Println("Subtracting function 1 - 2 is: ", sub(1,2))
	fmt.Println("Multiplying function 1 * 2 is: ", mul(1,2))
	fmt.Println("Dividing function 1 / 2 is: ", div(1,2))
	fmt.Println("********************************************")
}

func add(augend, addend int) int {
    return augend + addend
}


func sub(minuend, subtrahend int) int {
	return minuend - subtrahend
}

func mul(multiplier, multiplicand int) int {
	return multiplier * multiplicand
}

func div(numerator, denominator int) int {
	return numerator / denominator
}
