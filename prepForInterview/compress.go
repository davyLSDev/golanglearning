package main

/*	- compress a given string which only contains (a-z)/(A-Z)
	- input string ="aabcccccaaa" . output string = "a2b1c5a3"
	- if the output string >= input string, then don't "compress"
*/

import (
	// need to work on modules, or packages ... later	"../getline"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// this is just a wrapper to test out the function
func main() {
	promptA := "Enter a string to compress"
	lineA := Fetch(promptA)
	fmt.Println("The result of compressing your input string is >", compress(lineA))
}

func compress(initialString string) (finalString string) {
	finalString = ""
	initialLength := len(initialString)
	var currentCharacter string
	previousCharacter := ""
	characterCount := 0 // the number of characters for any given character sequence in the initial string

	for characterIdx := 0; characterIdx < initialLength; characterIdx++ {
		currentCharacter = string(initialString[characterIdx])

		// for first character in the string
		if previousCharacter == "" {
			previousCharacter = currentCharacter
		}

		if currentCharacter == previousCharacter {
			characterCount++

			if characterIdx+1 == initialLength { // ensure string is printed out
				finalString = finalString + string(currentCharacter) + strconv.Itoa(characterCount) // proper way to convert integer to string
			}

		} else {
			finalString = finalString + string(previousCharacter) + strconv.Itoa(characterCount)
			characterCount = 1
		}
		if initialLength <= len(finalString) { // compression otherwise would produce longer string
			finalString = initialString
			break
		}
		previousCharacter = currentCharacter

	}

	return finalString
}

func Fetch(prompt string) (lineOfCharacters string) {

	fmt.Println(prompt)
	/* Note to include whitespace when fetching input, you need to use bufio.Scanner, see
			   https://stackoverflow.com/questions/43843477/scanln-in-golang-doesnt-accept-whitespace
	   	   	   	fmt.Scanln(&aString)
	*/
	keyScanner := bufio.NewScanner(os.Stdin)
	if keyScanner.Scan() {
		lineOfCharacters = keyScanner.Text()
	}

	return lineOfCharacters
}
