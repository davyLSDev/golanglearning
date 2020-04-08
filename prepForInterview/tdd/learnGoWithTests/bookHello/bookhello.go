package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

// Hello is the main function
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Dawson", ""))
}
