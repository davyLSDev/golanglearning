package main

import (
	"bufio"
	"fmt"
	"os"
)

// this is just a wrapper to test out the function
func main() {
	var aString string
	var bString string

	// fmt.Println("Enter a string.")

	// keyScanner := bufio.NewScanner(os.Stdin)
	// if keyScanner.Scan() {
	// 	aString = keyScanner.Text()
	// }
	aString = getString("Enter the first string")
	bString = getString("Enter the next string")
	if oneAway(aString, bString) {
		fmt.Println("The strings are one away or less")
	} else {
		fmt.Println("The strings are more than one away")
	}

}

/* here is where the real work happens
   Edits are:
   1) insert a character
   2) remove a character
   3) replacea a character
*/
func oneAway(firstString, secondString string) bool {
	var isOneAway bool
	isOneAway = false
	firstStringLength := len(firstString)
	secondStringLength := len(secondString)
	// clearly the strings are more than one edit away
	if abs(firstStringLength-secondStringLength) > 1 {
		return isOneAway
	}
	// clearly the strings are one or less than one edit away
	if firstString == secondString {
		isOneAway = true
	}
	return isOneAway
}

func abs(number int) int {
	if number < 0 {
		return -number
	} else {
		return number
	}
}

func getString(prompt string) string {
	var stringEntered string
	fmt.Println(prompt)

	keyScanner := bufio.NewScanner(os.Stdin)
	if keyScanner.Scan() {
		stringEntered = keyScanner.Text()
	}
	return stringEntered
}
