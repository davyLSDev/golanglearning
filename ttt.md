Notes on how to program ttt program

This package has a really dumb ai bot which randomly picks the
next place to play after the user plays

## Some thoughts about how to do this

* MVP, keep it simple: player starts as x (https://en.wikipedia.org/wiki/Tic-tac-toe)
* using ANSI/VT100 Terminal Control Escape Sequences (http://www.termsys.demon.co.uk/vtansi.htm)
* probably predefine clearscreen, setcursor as character sequences
* 3x3 array to keep the game board information
* conversions between 1-9, cursor positions, and game array postions
* random function to generate a random number
* time in seconds to seed the random number generator
* slice to keep the random numbers left over from the playing board

## Some functions that are likely needed

* setup new game
* display
* get keyboard input from player, X
* display
* check for a win / scratch, MVP use brute force
* [keep track of plays for more elegant/efficient check for score]
* generate random number (1-9) for the squares available & play O
* check for win /scratch
* game over when there is a win or scratch
* announce a winner
* [fancier game add a line strike through for the win]

## Display of game board

```
	  |   |          7 | 8 | 9
	---------        ---------
	  |   |          4 | 5 | 6
	---------        ---------
	  |   |          1 | 2 | 3

	Your turn, choose a number to place your X
```

* an example format for working with the screen: fmt.Println("\033[2J\033[25;10HHello")
* \033 is <ESC>
* \033[2J is clear screen
* \033[25;10H is set cursor home to row 25, and column 10