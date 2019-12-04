// checkEquivalentKeypresses is a function that takes an array containing
// two strings where each string represents keypresses separated by commas.
// A keypress can be either a printable character or a backspace
// (represented by -B).
// The function determines if the two strings of keypresses are equivalent.

// Produce a printable string from such a string of keypresses by having
// backspaces erase one preceding character. Consider two strings of
// keypresses equivalent if they produce the same printable string.

// Example datasets:

// checkEquivalentKeypresses(["a,b,c,d", "a,b,c,c,-B,d"]) // true
// checkEquivalentKeypresses(["-B,-B,-B,c,c", "c,c"]) // true
// checkEquivalentKeypresses(["", "a,-B,-B,a,-B,a,b,c,c,c,d"]) // false

package main

import (
	"fmt"
	"strings"
)

func main() {
	datasets := [][]string{
		{"a,b,c,d", "a,b,c,c,-B,d"},
		{"-B,-B,-B,c,c", "c,c"},
		{"", "a,-B,-B,a,-B,a,b,c,c,c,d"},
	}

	for _, dataset := range datasets {
		test := checkEquivalentKeypresses(dataset)
		if test {
			fmt.Println("The two strings are equivalent")
		} else {
			fmt.Println("The two strings are not equivalent")
		}
	}
}

func checkEquivalentKeypresses(stringsToTest []string) (status bool) {
	stringA := editString(stringsToTest[0])
	stringB := editString(stringsToTest[1])
	status = false
	fmt.Println("stringA is ", stringA)
	fmt.Println("stringB is ", stringB)
	if stringA == stringB {
		status = true
	}
	return status
}

func editString(startString string) (editedString string) {
	var key string
	separator := ","
	backspace := "-B"
	sliceSequence := strings.Split(startString, separator)

	for _, keystroke := range sliceSequence {
		key = string(keystroke)
		editedStringLength := len(editedString)
		if key != backspace {
			editedString = editedString + key
		} else {
			if editedStringLength != 0 {
				editedString = editedString[:editedStringLength-1]
			}
		}
	}
	return editedString
}
