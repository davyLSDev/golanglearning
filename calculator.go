package main

import (
	// a great way to remember to import a package before using it       _ "fmt"

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// names for operands here: https://math.stackexchange.com/questions/975541/what-are-the-formal-names-of-operands-and-results-for-basic-operations

func main() {
	helptext := `NAME
	calculator
Usage
	./calculator [mathematical operation] operand1 operand2
Description
	Return the requested calculation result to standard output

AUTHOR
	Written by Dawson Tennant
REPORTING BUGS
	Are you kidding?
COPYRIGHT
	Copyright (c) 2019 Free Software Foundation, Inc.  License GPLv3+: GNU
	GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
	This is free software: you are free  to  change  and  redistribute  it.
	here is NO WARRANTY, to the extent permitted by law.
SEE ALSO
	Nothing to see here, move along now.
`
	switch argCount := len(os.Args[1:]); argCount {
	case 0:
		fmt.Println("No parameters given, assuming interractive mode")
		fmt.Println("For now check out sub 4 5, which results in: ", sub(4, 5))
		fmt.Println("Now we will print out what you type")
		interactive()
	case 1:
		fmt.Println("Assume interractive calculator then.")
		fmt.Println("For testing purposes, add 1 2, which results in: ", add(1, 2))
	case 2:
		fmt.Println("The calculator needs another operand, enter that now please.")
	case 3:
		fmt.Println("Much better, now we need to check for parameter sanity!")
		operation := os.Args[2]

		fmt.Println("There are ", argCount, "arguments entered")
		fmt.Println("This is a very simplisitic way to do command parsing")
		fmt.Println("You want to do the following arithmetic operation: ", operation)
		fmt.Println("For integer arithmetic")
		fmt.Println("Adding function 1 + 2 is: ", add(1, 2))
		fmt.Println("Subtracting function 1 - 2 is: ", sub(1, 2))
		fmt.Println("Multiplying function 1 * 2 is: ", mul(1, 2))
		fmt.Println("Dividing function 1 / 2 is: ", div(1, 2))
		fmt.Println("********************************************")
	default:
		fmt.Println(helptext)
	}
}

func interactive() {
	var number1, number2 float64
	dataentry := bufio.NewReader(os.Stdin)
	info, _ := dataentry.ReadString('\n')
	fmt.Println("This is what you typed: ", "\n", info)
	fmt.Println("Please enter an number")
	number1 = fetchnumber()
	fmt.Println("Please enter the second number")
	number2 = fetchnumber()
	fmt.Println("These are the two numbers you entered ", number1, number2)
}

func fetchnumber() float64 {
	var number string
	getdata := bufio.NewReader(os.Stdin)
	data, _ := getdata.ReadString('\n')
	number = strings.TrimSuffix(data, "\n")
	fnumber, _ := strconv.ParseFloat(number, 64)
	return fnumber
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

func help() {
	fmt.Println("")
}