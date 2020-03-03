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
	//	lineA := getline.Fetch(promptA)
	lineA := Fetch(promptA)
	fmt.Println(lineA)
	fmt.Println("The result of compressing your input string is >", compress(lineA))
}

func compress(initialString string) (finalString string) {
	initialLength := len(initialString)
	var currentCharacter, lastCharacter string
	characterCount := 1 // duplicate characters

	for characterIdx := 0; characterIdx < initialLength; characterIdx++ {
		currentCharacter = string(initialString[characterIdx])

		if currentCharacter == lastCharacter {
			characterCount++
		} else {
			finalString = finalString + string(currentCharacter) + strconv.Itoa(characterCount)
			fmt.Println("Debug, characterCount is", characterCount)
			fmt.Println("Attmept to stringify characterCount is", strconv.Itoa(characterCount))
			//			fmt.Println("This is the correct way to convert integer to string", strconv.Itoa(characterCount))
			characterCount = 0
		}
		lastCharacter = currentCharacter

		if initialLength == len(finalString) {
			finalString = initialString
			break
		}

	}

	/*	if finalLength+1 == initialLength {
			finalString = initialString
		}
	*/
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
