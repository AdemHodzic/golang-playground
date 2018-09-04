package main

import "testing"

func TestFactorial(t *testing.T) {

	var cases = []struct {
		input int
		expected int
	} {
		{5,120},
		{10,3628800},
		{0,1}}

	for _, testCase := range cases {
		if Factorialize(testCase.input) != testCase.expected {
			panic("Test failed...")
		}
	}

}