package main

import (
	"testing"
	"fmt"
)

func TestWarmer(t *testing.T) {

	var tests = []struct {
		input []int
		output []int
	} {
		{[]int {73, 74, 75, 71, 69, 72, 76, 73}, []int {1, 1, 4, 2, 1, 1, 0, 0}},
	}

	for index, test := range tests {
		if !equal(Warmer(test.input), test.output) {
			fmt.Printf("FAILED TEST #%d\n", index + 1)
			fmt.Printf("GOT: %v\n", Warmer(test.input))
			t.Fail()
		}

	}

}

func equal(first []int, second []int) bool {

	for i, v := range first {
		if v != second[i] {
			return false
		}
	}

	return true

}