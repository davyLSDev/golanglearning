// Q9.1 This simple LIFO stack can hold a fixed number of ints.
//   Functions:
//     push – put something on the stack
//     pop  – retrieve something from the stack.

// Ex of simple stack
//				i++
//		******* <-- push(k)
//	  i	*  k  *		  i--
//		******* pop() --> k
//		*  l  *
//		*******
//	  0	*  m  *
//		*******

// Q9.2 String method converts the stack to a string representation,
// so as to be able to print the stack using:
// fmt.Printf("My stack %v\n", stack)

// The stack in the figure could be represented as: [0:m] [1:l] [2:k]
// from LearningGo book pages 32,33

// Useful reference to understand the difference between functions and methods in Go:
// https://golangbot.com/methods/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is going to test the stack function")
}
