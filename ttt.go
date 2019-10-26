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

// func init() {
// 	restart()
// 	// fmt.Println("Do this first")
// 	//	return 1
// }

// func restart() {
// 	fmt.Println("Do this first")
// }

func main() {
	// Wait to implement restart()
	display(0, (' '))
	playerMark := byte('X')
	robotMark := byte('O')
	var location int

	for {
		_, err := fmt.Scanf("%d", &location)
		fmt.Println(err)
		mark := playerMark
		display(location-1, mark)
		mark = robotMark
		location := ai()
		display(location-1, mark)
	}
}

func display(placement int, character byte) {
	clear := "\033[2J"
	// x := byte('X')
	// y := byte('Y')
	row := 25
	column := 1
	message := []byte(`
  |   |          1 | 2 | 3
---------        ---------
  |   |          4 | 5 | 6
---------        ---------
  |   |          7 | 8 | 9

Your turn, choose a number to place your X, and press <ENTER>`)

	// index is the array to place X's and O's into the grid
	// remember the index starts at 0
	index := [9]int{1, 5, 9, 55, 59, 63, 109, 113, 117}
	message[index[placement]] = character
	output := strings.Join([]string{"\033[" + strconv.Itoa(row) + ";" + strconv.Itoa(column) + "H" + string(message)}, "+ ")
	fmt.Println(clear, output)
}

func ai() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(9)
}
