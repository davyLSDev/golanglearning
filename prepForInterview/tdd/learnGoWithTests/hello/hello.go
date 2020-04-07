package main

import (
	"fmt"
	"os"
)

var name string

const (
	cs = iota // Czech
	de        // German
	en
	es // Spanish
	fr // French
	ga // Irish
	it // Italian
	nl // Dutch
	pt // Portugeses
)

// according to https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
var languageLookup = []struct {
	language string
	greeting string
}{
	{"Czech", "Ahoj"},
	{"German", "Hallo"},
	{"English", "Hello"},
	{"Spanish", "Hola"},
	{"French", "Bonjour"},
	{"Irish", "Dia dhuit"},
	{"Italian", "Ciao"},
	{"Dutch", "Hallo"},
	{"Portugese", "Olá"},
}

// Greet takes a name string and a language code to determine an appropriate greeting in the language of choice
func Greet(name string, language int) (translated string) {
	var message string
	message = languageLookup[language].greeting
	translated = message + ", " + name
	return translated
}

const (
	programExitOk   = 0
	programExitFail = 1
)

func main() { os.Exit(Program()) }

// Program the testable main
func Program() (exitcode int) {
	var languageCode int
	exitcode = programExitOk
	name = "Dawson"
	fmt.Println(Greet(name, languageCode))
	return exitcode
}
