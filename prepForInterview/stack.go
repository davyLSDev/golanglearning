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
	"strconv"
)

var stack []int

func main() {
	size := 10 // stack has a fixed size of some kind
	var result, err string
	var value int

	showStack()
	for i := 10; i < 22; i++ {
		err := push(i, size)
		result = "Successfully pushed " + strconv.Itoa(i) + " onto the stack."
		printResults(result, err)
	}
	showStack()

	for j := 10; j < 22; j++ {
		err, value = pop()
		result = "The value popped of the stack is " + strconv.Itoa(value) + "."
		printResults(result, err)
	}
	showStack()
}

func showStack() {
	fmt.Println("This is my stack now", stack)
	return
}

func push(value, maximumSize int) (err string) {
	err = ""
	currentStackSize := len(stack)
	if currentStackSize == maximumSize {
		err = "Stack overflow!"
		// return err
	} else {
		stack = append(stack, value)
	}
	return err
}

func pop() (err string, value int) {
	err = ""
	if len(stack) <= 0 {
		err = "Stack underflow!"
		value = 0
	} else {
		value = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return err, value
}

func printResults(message, err string) {
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}
	return
}
