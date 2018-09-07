package main

import (
	"testing"
	"fmt"
)

func TestBinaryDistance(t *testing.T) {

	var tests = []struct{
		input int
		expected int
	} {
		{22,2},
		{5,2},
		{6,1},
		{8,0},
	}

	for index, test := range tests {
		if BinaryDistance(test.input) != test.expected {
			fmt.Printf("FAILED TEST #%v \n\t\tGOT: %v\n\t\tEXPECTED: %v\n", index + 1, BinaryDistance(test.input), test.expected)
			t.Fail()
		}
	}

}