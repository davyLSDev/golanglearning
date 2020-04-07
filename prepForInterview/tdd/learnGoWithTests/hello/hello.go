package main

import "fmt"

const en int = 1 // Spanish code
const es int = 2 // Spanish code

// Hello takes a name and a language string to determine an appropriate greeting
func Greet(name string, language int) (translated string) {
	var message string
	if language == en {
		message = "Hello"
	} else if language == es {
		message = "Hola"
	}
	translated = message + ", Dawson"
	return translated
}

func main() {
	var languageCode int
	fmt.Println(Greet("Dawson", languageCode))
}
