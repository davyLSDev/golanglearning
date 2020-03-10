package main

/* Simple linked list which allows for
(getting input from the user)
1. adding new nodes to the list
2. print out the list of nodes like HEAD: 12 -> 33 -> TAIL: 44
*/

import (
	"bufio"
	"fmt"
	"os"
)

const CLEAR = "\033[2J"
const PROMTP1 = CLEAR + `
**************************************
* Choose one of the following        *
*                                    *
* 0   DEBUG linked list printout     *
* 1   print out the linked list      *
* 2   insert new value to the list   *
*                                    *
**************************************`
const PROMPT2 = CLEAR + `
**************************************
*  Enter a value to add to the list  *
**************************************`

func main() {
	var choice, value string
	choice = Fetch(PROMTP1)
	value = Fetch(PROMPT2)
	fmt.Println("Your command was ", choice)
	fmt.Println("Your value was", value)
}

func Fetch(prompt string) (lineOfCharacters string) {

	fmt.Println(prompt)
	keyScanner := bufio.NewScanner(os.Stdin)
	if keyScanner.Scan() {
		lineOfCharacters = keyScanner.Text()
	}

	return lineOfCharacters
}
