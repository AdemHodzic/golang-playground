package main

import ("testing"
"fmt")

func TestInserter(t *testing.T) {
	var tests = []struct {
		arr []int
		number int
		expected int
	}{
		{[]int {10, 20, 30, 40, 50}, 35, 3},
		{[]int {10, 20, 30, 40, 50}, 30, 2},
		{[]int {40, 60}, 50, 1},
		{[]int {3, 10, 5}, 3, 0},
	}

	for index, test := range tests {
		if Inserter(test.arr,test.number) != test.expected {
			fmt.Printf("TEST FAILED #%v IN INSERTER TEST\n", index + 1)
			fmt.Printf("ARRAY: %v\n NUMBER: %d\n ", test.arr,test.number)
			t.Fail()
		}
	}
}


