package end

import ("testing"
)

func TestEndsWith(t *testing.T) {

	var cases = []struct {
		original string
		end string
		expected bool
	} {
		{"Bastian", "n", true},
		{"Congratulation", "on", true},
		{"Connor", "n", false},
	}

	for _, test := range cases {
		if EndsWith(test.original, test.end) != test.expected {
			panic("TEST FAILED...")
			
		}
	}

}