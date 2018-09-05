package end

import "testing"

func TruncateTest(t *testing.T) {
	var tests = []struct{
		str string
		len int
		expected string
	} {
		{"A-tisket a-tasket A green and yellow basket", 8, "A-tisket..."},
		{"A-",1,"A..."},
	}

	for _, test := range tests {
		if Truncate(test.str,test.len) != test.expected {
			panic("TEST FAILED...")
		}
	}
}