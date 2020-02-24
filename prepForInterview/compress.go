package main

import (
	// need to work on modules, or packages ... later	"../getline"
	"bufio"
	"fmt"
	"os"
)

// this is just a wrapper to test out the function
func main() {
	promptA := "Enter a string"
	//	lineA := getline.Fetch(promptA)
	lineA := Fetch(promptA)
	fmt.Println(lineA)
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
