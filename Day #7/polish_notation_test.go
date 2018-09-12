package main 

import (
	"testing"
	"fmt"
)

func TestPolishNotatin(t *testing.T) {

	var tests = []struct {
		input []string
		expected int
	} {
		{[]string {"2", "1", "+", "3", "*"}, 9},
		{[]string {"4", "13", "5", "/", "+"}, 6},
		{[]string {"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for index, test := range tests {

		if PolishNotation(test.input) != test.expected {
			fmt.Printf("FAILED TEST #%d\n", index + 1)
			fmt.Printf("GOT: %v\n", PolishNotation(test.input))
			t.Fail()
		}

	}

}