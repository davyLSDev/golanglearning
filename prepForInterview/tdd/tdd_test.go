/* Tutorial(?) https://codingislove.com/test-driven-development-golang/
Notes: missing front matter, it is just code snippets

I did need to add this package thus: go get github.com/stretchr/testify

I had to add "package main"
When I saved this file, some intelligence added the import lines.

Oh, of course, this is the file that you start with!
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiples(t *testing.T) {
	// Return multiples of 3
	assert.Equal(t, []int{3}, Multiples(1, 5))

	// Return multiples of 5 also
	assert.Equal(t, []int{5, 6, 9, 10}, Multiples(4, 11))

	// Return empty slice if multiples are not found
	assert.Empty(t, []int{}, Multiples(1, 3))
}
