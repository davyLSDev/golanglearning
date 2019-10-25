/* This package has a really dumb ai bot which randomly picks the
next place to play after the user plays */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	display()
}

func display() {
	clear := "\033[2J"
	row := 25
	column := 1
	message := `
  |   |          7 | 8 | 9
---------        ---------
  |   |          4 | 5 | 6
---------        ---------
  |   |          1 | 2 | 3

Your turn, choose a number to place your X.
	`
	output := strings.Join([]string{"\033[" + strconv.Itoa(row) + ";" + strconv.Itoa(column) + "H" + message}, "+ ")
	fmt.Println(clear, output)
}
