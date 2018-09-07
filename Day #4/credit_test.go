package main

import (
	"testing"
	"fmt"
	"strings"
)

func TestCheckCard(t *testing.T) {

	var tests = []struct {
		input int
		expected string
	} {
		{378282246355005, "American Express"},
		{6176292929, "Invalid"},
	}

	for index, test := range tests {
		if !strings.EqualFold(CheckCard(test.input), test.expected) {
			fmt.Printf("FAILED TEST #%v \n\t\tGOT: %v\n\t\tEXPECTED: %v\n", index + 1, CheckCard(test.input), test.expected)
			t.Fail()
		}
	}

}