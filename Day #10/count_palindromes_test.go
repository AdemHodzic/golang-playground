package main 

import (
	"testing"
	"fmt"
)

func TestCountPalindromes(t *testing.T) {

	var tests = []struct {
		input string
		expected int
	} {
		{"abc", 5},
		{"aaa", 6},
	}

	for index, test := range tests {

		if CountPalindromes(test.input) != test.expected{
			fmt.Printf("FAILED AT TEST #%d\n", index + 1)
			fmt.Printf("GOT %v\n", CountPalindromes(test.input))
			t.Fail()
		}

	}

}