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
	"time"
)

const FINISH = true
const AGAIN = false
const PAUSE = 1 * time.Second
const NOPAUSE = 0 * time.Second
const CLEAR = "\033[2J"
const PROMTP1 = CLEAR + `
**************************************
* Choose one of the following        *
*                                    *
* 0   DEBUG linked list printout     *
* 1   print out the linked list      *
* 2   add a new value to the list    *
* 3   exit                           *
*                                    *
**************************************`
const PROMPT2 = CLEAR + `
**************************************
*  Enter a value to add to the list  *
**************************************`
const ERROR1 = CLEAR + `
**************************************
*  Please enter a number between 0-2!*
**************************************`
const ACTION0 = CLEAR + "DEBUG printout of list"
const ACTION1 = CLEAR + "Print out the linked list"
const ACTION2 = CLEAR + "Fetch a value to add to the list"
const ACTION3 = CLEAR + "Exit now"

type Node struct {
	data string
	next *Node
}

type List struct {
	int length
	head *Node
}

func main() {
	var action, choice, value string
	loopState := AGAIN
	pauseTime := PAUSE

	for {
		choice = Fetch(PROMTP1)
		switch choice {
		case "0":
			action = ACTION0
		case "1":
			action = ACTION1
		case "2":
			action = ACTION2
			value = Fetch(PROMPT2)
			fmt.Println(CLEAR+"Your value was", value, "it will now be added to the list")
			time.Sleep(pauseTime)
		case "3":
			action = ACTION3
			loopState = FINISH
		default:
			action = ERROR1
		}

		if loopState == FINISH {
			// break
			fmt.Println("The loop should end now")
			time.Sleep(pauseTime)
			break
		}
	}

	fmt.Println(action)
	time.Sleep(pauseTime)
}

func Fetch(prompt string) (lineOfCharacters string) {

	fmt.Println(prompt)
	keyScanner := bufio.NewScanner(os.Stdin)
	if keyScanner.Scan() {
		lineOfCharacters = keyScanner.Text()
	}

	return lineOfCharacters
}
