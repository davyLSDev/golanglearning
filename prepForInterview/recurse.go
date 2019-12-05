// recursion ... try to understand this

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Use the recursive function with starting value of 0")
	rec(0)
}

func rec(i int) {
	fmt.Println("First line of function and the value of i is ", i)
	if i == 10 {
		fmt.Println("The value of i is 10 and the function returns")
		return
	} else {
		fmt.Println("The value of i is not 10 yet")
	}
	rec(i + 1)
	fmt.Printf("%d ", i)
	fmt.Println("This is i after return ", i)
}
