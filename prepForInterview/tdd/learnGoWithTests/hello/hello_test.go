package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Dawson", es)
	want := "Hola, Dawson"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	got = Greet("Dawson", en)
	want = "Hello, Dawson"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
