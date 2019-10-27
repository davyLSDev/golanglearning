package main

import (
	"fmt"
)

func main() {

	lines := [8][3]int{
		{0, 1, 2}, // row one
		{3, 4, 5}, // row two
		{6, 7, 8}, // row three
		{0, 3, 6}, // column one
		{1, 4, 7}, // column two
		{2, 5, 8}, // column three
		{0, 4, 8}, // diagonol one
		{2, 4, 6}} // diagonal two

	for i, x := range lines {
		for j, _ := range x {
			fmt.Println(lines[i][j])
		}
	}
}
