package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThings(t *testing.T) {
	// equality
	assert.Equal(t, 123, 123, "They should be equal.")
	// inequality
	assert.NotEqual(t, 123, 456, "They should not be equal.")

	// This next bit doesn't work correctly.
	// assert for nil ... good for errors according to the author
	// but without defining object, you will get "undefined object ... Error: Tests failed" because it won't compile
	// assert.Nil(t, object)
	// if assert.NotNil(t, object) {
	// 	assert.Equal(t, "", object.Value)
	// }
	assert.NotNil(t, object)
}
