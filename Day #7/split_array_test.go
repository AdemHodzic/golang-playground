package main 

import (
	"testing"
	"fmt"
	
)

func TestSplit(t *testing.T) {

	var tests = []struct {
		input []int 
		expected bool
	} {
		{[]int {1,2,3,4,5,6,7,8}, true},
	}

	for index, test := range tests {
		if Split(test.input) != test.expected {
			fmt.Printf("FAILED TEST #%d\n", index + 1)
			fmt.Printf("GOT: %v\n", Split(test.input))
			t.Fail()
		}

	}

}