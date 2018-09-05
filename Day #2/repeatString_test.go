package end

import ("testing")

func TestRepeatString(t *testing.T) {
	var tests = []struct {
		str string
		times int
		expected string
	} {
		{"*",3,"***"},
		{"abc", 3, "abcabcabc"},
		{"abc", 1, "abc"},
		{"abc", -2, ""},
	}

	for _, test := range tests {
		if RepeatString(test.str, test.times) != test.expected {
			panic(test.str)
		}
	}
}