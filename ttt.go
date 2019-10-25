/* This package has a really dumb ai bot which randomly picks the
next place to play after the user plays */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
	
	Some thoughts about how to do this

	- MVP, keep it simple: player starts as x (https://en.wikipedia.org/wiki/Tic-tac-toe)
	- 3x3 array to keep the game board information
	- random function to generate a random number
	- time in seconds to seed the random number generator
	- slice to keep the random numbers left over from the playing board
	
	Some functions that are likely needed
	
	- setup new game
	- display
	- get keyboard input from player, X
	- display
	- check for a win / scratch, MVP use brute force
	- [keep track of plays for more elegant/efficient check for score]
	- generate random number (1-9) for the squares available & play O
	- check for win /scratch
	- game over when there is a win or scratch
	- announce a winner
	- [fancier game add a line strike through for the win]

	Display of game board

	  |   |          7 | 8 | 9
	---------        ---------
	  |   |          4 | 5 | 6
	---------        ---------
	  |   |          1 | 2 | 3

	Your turn, choose a number to place your X
	`)
}
