// https://tutorialedge.net/golang/improving-your-tests-with-testify-go/
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
}
