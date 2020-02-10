package main
/*
	Replace all spaces in a string with %20
	- the string holds enough spaces at the end for additional characters
	- the true length of the string is given
	Ex: given "Mr John Smith    ",13 produce "Mr%20John%20Smith"
*/

import (
	"fmt"
)

func main() {
	initialStr := "Jo Blow  "
	trueStrLength := 7

	fmt.Println("Initial string is", initialStr)
	urlifiedStr := urlify(initialStr, trueStrLength)
	fmt.Println("URLified version of the string is", urlifiedStr)
}

func urlify(input string, length int) string {
	urlSequence := "%20"
	urlifiedStr := ""
	var nextChars string

	for character := 0; character < length; character++ { // if you do character == loop doesn't do anything
		nextChars = input[character:character+1]
		if input[character:character+1] == " " {
			nextChars = urlSequence
		}
		fmt.Println("next characters is", nextChars)
		urlifiedStr = urlifiedStr + nextChars
		fmt.Println("This is loop number", character)
	}
	return urlifiedStr
}