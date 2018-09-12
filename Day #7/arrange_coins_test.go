package main 

import (
	"testing"
	"fmt"
)

func TestArrangeCoins(t *testing.T) {

	var tests = []struct {
		input int 
		expected int 
	} {
		{5,2},
		{8,3},
	}

	for index, test := range tests {
		if ArrangeCoins(test.input) != test.expected {
			fmt.Printf("FAILED TEST #%d\n", index + 1)
			fmt.Printf("GOT: %v\n", ArrangeCoins(test.input))
			t.Fail()
		}

	}

}