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

// Useful references to understand the difference between functions and methods in Go:
// https://golangbot.com/methods/, and even closer to what I need: https://www.callicoder.com/golang-methods-tutorial/
// https://tour.golang.org/methods/1
// on maps: https://appdividend.com/2019/05/12/golang-maps-tutorial-with-examples-maps-in-go-explained/

// Actually I don't really see the point of using a map for this, a slice seems much more normal.

// Maybe what we need is stackInitialize, push, pull, stack string method which work on a type "stack"
// that I define.

package main

import (
	"fmt"
)

func main() {
	const stackSize = 10 // stack has a fixed size of some kind
	var myStack []int
	fmt.Println("This is going to test the stack function / method")
	push(myStack, 55, stackSize)
	showStack(myStack)
}

func showStack(stackNow []int) {
	fmt.Println("This is my stack now", stackNow)
	return
}

func push(stackNow []int, valueToPush int, stackSize int) (err string) {
	err = ""
	stackPointer := len(stackNow)
	if stackPointer >= stackSize {
		err = "Stack overflow"
	} else {
		fmt.Println(stackPointer)
		// stackNow[stackPointer] = valueToPush
	}
	return err
}
