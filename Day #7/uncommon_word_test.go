package main 

import (
	"testing"
	"fmt"
	"strings"
)

func TestUncommonWords(t *testing.T) {

	var tests = []struct {
		first string
		second string
		expected []string
	} {
		{"this apple is sweet","this apple is sour", []string {"sweet","sour"}},
		{"apple apple", "banana", []string {"banana"}},
	}

	for index, test := range tests {
		if !equalStrings(UncommonWords(test.first, test.second), test.expected) {
			fmt.Printf("FAILED TEST #%d\n", index + 1)
			fmt.Printf("GOT: %v\n", UncommonWords(test.first, test.second))
			t.Fail()
		}

	}

}


func equalStrings(first []string, second []string ) bool {

	for index, value := range first {

		if !strings.EqualFold(value, second[index]) {
			return false
		}

	}

	return true

}
