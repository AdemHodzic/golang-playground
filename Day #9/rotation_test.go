package main

import (
	"fmt"
	"testing"
)

func TestCheckRotation(t *testing.T) {

	var tests = []struct {
		first    string
		second   string
		expected bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
	}

	for index, test := range tests {
		if CheckRotation(test.first, test.second) != test.expected {
			fmt.Printf("FAILED AT TEST #%v\n", index+1)
			fmt.Printf("GOT %v\n", CheckRotation(test.first, test.second))
			t.Fail()
		}
	}

}
