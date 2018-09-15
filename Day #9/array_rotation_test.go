package main 

import (
	"testing"
	"fmt"
)

func TestBestRotation(t *testing.T) {

	var tests = []struct {
		arr []int 
		expected int
	} {
		{[]int {2, 3, 1, 4, 0},3},
		{[]int {1, 3, 0, 2, 4}, 0},
	}

	for index,test := range tests {
		if BestRotation(test.arr) != test.expected {
			fmt.Printf("FAILED AT TEST #%d\n", index + 1)
			fmt.Printf("GOT %v\n", BestRotation(test.arr))
			t.Fail()
		}
	}

}
