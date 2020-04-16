package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumALL(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	/* Go does not let you use equality operators with slices. The tutorial author suggests
	that I could write a function to iterate over each got and want slice and check their
	values, but for convenience, we can use "reflect.DeepEqual" which is useful for seeing
	if any two variables are the same.
	if got != want {
	*/

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
