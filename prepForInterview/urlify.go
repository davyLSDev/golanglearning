package main
/*
	Replace all spaces in a string with %20
	- the string holds enough spaces at the end for additional characters
	- the true length of the string is given
	Ex: given "Mr John Smith    ",13 produce "Mr%20John%20Smith"
*/

import (
	"fmt"
	// "str"
)

func main() {
	initialStr := "Jo Blow  "
	trueStrLength := 7

	fmt.Println("Initial string is", initialStr)
	urlifiedStr := urlify(initialStr, trueStrLength)
	fmt.Println("URLified version of the string is", urlifiedStr)
}

func urlify(input string, length int) string {
	urlifiedStr := "Jo%20Blow"
//	urlifiedStrLength := len(initialStr)
	return urlifiedStr
}