package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	/* Notice the slight difference between how we called assert.Equal() in this example compared to the
	previous example. We’ve initialized assert using assert.New(t) and we are now able to call
	assert.Equal() multiple times, just passing in the input and the expected values as opposed to having
	to pass t in as our first parameter every time. This isn’t a big deal, but it certainly helps to make
	our tests look cleaner
	*/
	assert := assert.New(t)

	// Here we can use a table to exercise the Calculate function
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
