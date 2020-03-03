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

/* here is where the real work happens
the headChar at headCharIdx is the current character to test against all remaining
characters in the given string
*/
func isUnique(inputString string) bool {
	unique := true
	stringLength := len(inputString)
//	headCharIdx := 0
//	fmt.Println("This string is",stringLength,"characters long")
	for headCharIdx := 0; headCharIdx < (stringLength-1); headCharIdx++ {
		for charIdx := (headCharIdx+1); charIdx < stringLength; charIdx++ {
			if inputString[headCharIdx] == inputString[charIdx] {
				unique = false
				break
			}
		}
	}
	return unique
}
