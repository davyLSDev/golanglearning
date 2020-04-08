package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

// Hello is the main function
func Hello(name, language string) string {
	// prefix := englishHelloPrefix

	if name == "" {
		name = "World"
	}

	// switch language {
	// case french:
	// 	prefix = frenchHelloPrefix
	// case spanish:
	// 	prefix = spanishHelloPrefix
	// }

	// return prefix + name
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	prefix = englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Dawson", ""))
}
