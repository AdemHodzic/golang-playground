package main

import (
	"testing"
	"fmt"
)

func TestPermutationsPower(t *testing.T) {

	var tests = []struct {
		input int
		expected bool
	} {
		{1,true},
		{10, false},
		{61, true},
		{24, false},
		{46, true},
	}

	for index, test := range tests {
		if PermutationsPower(test.input) != test.expected {
			fmt.Printf("FAILED TEST #%v \n\t\tGOT: %v\n\t\tEXPECTED: %v\n", index + 1, PermutationsPower(test.input), test.expected)
			t.Fail()
		}
	}


}