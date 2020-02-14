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
	fmt.Println("This is your first string", aString)
	fmt.Println("This is your second string", bString)
}

/* here is where the real work happens


func oneAway(firstString, secondString string) {

}
*/

func getString(prompt string) string {
	var stringEntered string
	fmt.Println(prompt)

	keyScanner := bufio.NewScanner(os.Stdin)
	if keyScanner.Scan() {
		stringEntered = keyScanner.Text()
	}
	return stringEntered
}
