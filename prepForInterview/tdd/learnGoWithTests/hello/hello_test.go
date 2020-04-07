package main

import "testing"

// according to https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
var allLanguagesTest = []struct {
	languageCode       int
	translatedGreeting string
}{
	{cs, "Ahoj"},  // Czech
	{de, "Hallo"}, // German
	{en, "Hello"},
	{es, "Hola"},      // Spanish
	{fr, "Bonjour"},   // French
	{ga, "Dia dhuit"}, // Irish
	{it, "Ciao"},      // Italian
	{nl, "Hallo"},     // Dutch
	{pt, "Olá"},       // Portugese
}

func TestGreet(t *testing.T) {
	for _, tt := range allLanguagesTest {
		got := Greet(name, tt.languageCode)
		want := tt.translatedGreeting + ", " + name
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}

func TestProgram(tp *testing.T) {
	got := Program()
	want := programExitOk
	if got != want {
		tp.Errorf("got %q want %q", got, want)
	}
}
