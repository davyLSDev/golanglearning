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
)

func main() {
	datasets:= [][]string{
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
	if stringA == stringB {
		status = true
	}
	return status
}

func editString(startSequence string) (editedSequence string) {
	var buildSequence, key string
	//length := len(startSequence)
	backspace := "-B"
	for _,keystroke := range startSequence {
		key = string(keystroke)
		fmt.Println("Key ", key)
		buildSequence = buildSequence + key
	}
	fmt.Println("The built sequence is ", buildSequence)
	editedSequence = startSequence
	if backspace != "" {
		fmt.Println("simple if")
	}
	fmt.Println("Still no editing on the keystroke sequence")
	return editedSequence
}
