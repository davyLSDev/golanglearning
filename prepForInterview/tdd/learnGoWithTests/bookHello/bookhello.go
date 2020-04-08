package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello is the main function
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Dawson"))
}
