// tutorial from https://tutorialedge.net/golang/improving-your-tests-with-testify-go/
package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
