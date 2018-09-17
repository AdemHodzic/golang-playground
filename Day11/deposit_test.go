package day11

import (
	"fmt"
	"testing"
)

func TestCalculateDeposit(t *testing.T) {

	tests := []struct {
		deposit  float32
		rate     int
		max      float32
		expected int
	}{
		{100, 20, 170, 3},
		{100, 1, 101, 1},
		{50, 25, 100, 4},
	}
	for index, test := range tests {
		if CalculateDeposit(test.deposit, test.rate, test.max) != test.expected {
			fmt.Printf("FAILED AT TEST #%d\n", index+1)
			fmt.Printf("GOT %v\n", CalculateDeposit(test.deposit, test.rate, test.max))
			t.Fail()
		}
	}
}
