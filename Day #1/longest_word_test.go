package main

import "testing"

func TestingLongestWord(t *testing.T) {
	var cases = []struct{
		input string
		expected int
	}{
		{"The quick brown fox jumped over the lazy dog", 6},
		{"May the force be with you", 5},
		{"Google do a barrel roll", 6},
		{"What is the average airspeed velocity of an unladen swallow",8},
		{"What if we try a super-long word such as otorhinolaryngology", 19},
	}

	for _, test := range cases {
		if LongestWord(test.input) != test.expected {
			panic("Test failed...")
		}
	}
}