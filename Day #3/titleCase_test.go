package main

import ("testing"
"fmt")

func TestTitleCase(t *testing.T) {
	var tests = []struct {
		input string
		expected string
	} {
		{"I'm a little tea pot","I'm A Little Tea Pot"},
		{"sHoRt AnD sToUt", "Short And Stout"},
		{"HERE IS MY HANDLE HERE IS MY SPOUT", "Here Is My Handle Here Is My Spout"},
	}

	for index,test := range tests {
		if TitleCase(test.input) != test.expected {
			fmt.Printf("Failed at test number %b", index + 1)
			t.Fail()
		}
	}
}