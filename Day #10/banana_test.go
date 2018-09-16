package main 

import (
	"testing"
	"fmt"
)

func TestBananaSpeed(t *testing.T) {

	var tests = []struct {
		piles []int 
		hours int
		expected int
	} {
		{[]int {3,6,7,11}, 8, 4}, 
		{[]int {30,11,23,4,20}, 5, 30}, 
		{[]int {30,11,23,4,20}, 6, 23}, 
	}

	for index, test := range tests {

		if BananaSpeed(test.piles, test.hours) != test.expected{
			fmt.Printf("FAILED AT TEST #%d\n", index + 1)
			fmt.Printf("GOT %v\n", BananaSpeed(test.piles, test.hours))
			t.Fail()
		}

	}

}