package main

import (
	"strings"
)

func CheckRotation(first string, second string) bool {

	if len(first) != len(second) {
		return false
	}

	maximum := len(first)

	for i := 0; i < maximum; i++ {

		var b strings.Builder
		b.Write([]byte(first[1:]))
		b.WriteByte(first[0])

		if strings.EqualFold(b.String(), second) {
			return true
		} else {
			first = b.String()
		}

	}

	return false
}
