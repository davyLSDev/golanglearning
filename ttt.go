/* This package has a really dumb ai bot which randomly picks the
next place to play after the user plays */
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var location int
	// player := 1
	// robot := 2
	playerMark := byte('X')
	robotMark := byte('O')
	board := []byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	playerPrompt := []byte(`Player's turn, choose a number to place your X, and press <ENTER>.`)
	robotPrompt := []byte(`Robot's turn.`)
	playerError := []byte(`Please enter a number between 1 to 9!`)
	redo := []byte(`That space is taken, enter a different number and press <ENTER>`)
	robotRedo := []byte(`Oh robot, try again!`)
	draw := []byte(`Game over, it's a draw!`)
	playerWins := []byte(`Game over, you have won.`)
	robotWins := []byte(`Game over, the robot won.`)

	mark := playerMark
	prompt := playerPrompt
	display(board, prompt)

	// redoPrompt := redo

	for {

		display(board, prompt)
		if mark != robotMark {
			// _, err := fmt.Scanf("%d", &location)
			// fmt.Println(err)
			location = playerInput()
			if location < 0 || location > 9 {
				prompt = playerError
				display(board, prompt)
			} else {
				if !checkPlace(board, location-1) {
					prompt = redo
					display(board, prompt)
				} else {
					board[location-1] = mark
					// display(board, prompt)
					mark = robotMark
					prompt = robotPrompt
				}
			}
		}

		if mark != playerMark {
			display(board, prompt)
			location := ai()
			if !checkPlace(board, location-1) {
				prompt = robotRedo
				display(board, prompt)
			} else {
				board[location-1] = mark
				mark = playerMark
				prompt = playerPrompt
				// display(board, prompt)
			}
		}

		scratch, win, winner := end(board)
		if scratch || win {
			if win {
				if winner == 1 {
					prompt = playerWins
				} else {
					prompt = robotWins
				}
			} else {
				scratch = true
				prompt = draw
			}
			display(board, prompt)
			break
		}
	}
}

func display(gameboard, status []byte) {
	clear := "\033[2J"
	row := 23
	column := 1
	message := []byte(`
  |   |          1 | 2 | 3
---------        ---------
  |   |          4 | 5 | 6
---------        ---------
  |   |          7 | 8 | 9

`)

	// index is the array to place X's and O's into the grid
	// remember the index starts at 0
	index := [9]int{1, 5, 9, 55, 59, 63, 109, 113, 117}

	for space := 0; space < 9; space++ {
		message[index[space]] = gameboard[space]
	}
	// message[index[placement]] = character
	output := strings.Join([]string{"\033[" + strconv.Itoa(row) + ";" + strconv.Itoa(column) + "H" + string(message) + string(status)}, "+ ")
	fmt.Println(clear, output)
	// fmt.Println(output)family@bigbird:~/projects/golang/golanglearning$
}

func ai() (aiPlay int) {
	rand.Seed(time.Now().UTC().UnixNano())
	testInteger := rand.Intn(9)
	aiPlay = testInteger + 1
	return
}

func playerInput() int {
	var selection int
	_, err := fmt.Scanf("%d", &selection)
	fmt.Println(err)
	return selection
}

func checkPlace(gameboard []byte, index int) (ok bool) {
	openPlace := byte(' ')
	ok = false
	if gameboard[index] == openPlace {
		ok = true
	}
	return
}

func end(cell []byte) (tie, win bool, champ int) {
	playerMark := byte('X')
	robotMark := byte('O')
	player := 1
	robot := 2
	var scratches int // when the sum of this is 8, the game is a tie!
	var playerMarkers int
	var robotMarkers int
	win = false
	tie = false

	row1 := []int{0, 1, 2}
	// row2 := []int{3, 4, 5}
	// row3 := []int{6, 7, 8}
	// col1 := []int{0, 3, 6}
	// col2 := []int{1, 4, 7}
	// col3 := []int{2, 5, 8}
	// diag1 := []int{0, 4, 8}
	// diag2 := []int{2, 4, 6}

	scratches = 0
	playerMarkers = 0
	robotMarkers = 0

	// this is repeatable code, need to consider how to refactor
	for _, rowIndex := range row1 {
		if cell[row1[rowIndex]] == playerMark {
			playerMarkers++
		}
		if cell[row1[rowIndex]] == robotMark {
			robotMarkers++
		}
		if playerMarkers == 3 || robotMarkers == 3 {
			win = true
			if playerMarkers == 3 {
				champ = player
			} else {
				champ = robot
			}
		}
		if playerMarkers > 0 && robotMarkers > 0 {
			scratches++
		}
	}
	return
}
