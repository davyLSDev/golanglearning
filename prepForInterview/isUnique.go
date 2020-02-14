package main

import (

	"fmt"
	"bufio"
	"os"
)

// this is just a wrapper to test out the function
func main() {
	var aString string

	fmt.Println("Enter a string to see if all characters are unique.")
/* Note to include whitespace when fetching input, you need to use bufio.Scanner, see
https://stackoverflow.com/questions/43843477/scanln-in-golang-doesnt-accept-whitespace
	fmt.Scanln(&aString)
*/
	keyScanner := bufio.NewScanner(os.Stdin)
	if keyScanner.Scan() {
		aString = keyScanner.Text()
	}
	fmt.Println("This is your string",aString)

	if isUnique(aString) {
		fmt.Println("All characters in", aString,"are unique.")
	} else {
		fmt.Println("Not all the characters in",aString,"are unique.")
	}
}

// here is where the real work happens
func isUnique(inputString string) bool {
	stringLength := len(inputString)
	fmt.Println("This string is",stringLength,"characters long")

	return false
}
