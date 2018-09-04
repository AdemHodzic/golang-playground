package main

import "strings"

func LongestWord(sentence string) int{
	arr := strings.Split(sentence, " ")
	max := 0

	for _, value := range arr {
		if len(value) > max {
			max = len(value)
		}
	}
	return max
}