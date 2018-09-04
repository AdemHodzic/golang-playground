package main

import "testing"

func TestingLargestFour(t *testing.T) {
	var cases = []struct {
		input [][]int
		expected []int
	} {
		{{{13, 27, 18, 26}, {4, 5, 1, 3}, {32, 35, 37, 39}, {1000, 1001, 857, 1}}, {27, 5, 39, 1001}},
	}

	for _, test := range cases {
		if LargestFour(test.input) != test.expected {
			panic("Test failed...")
		}
	}
}