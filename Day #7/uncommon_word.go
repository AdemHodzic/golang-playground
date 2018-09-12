package main 

import (
	"strings"
)

func UncommonWords(first string, second string) []string {


	var result []string

	first_arr := strings.Split(first, " ")
	second_arr := strings.Split(second, " ")

	for _, value := range first_arr {
		if strings.Count(first, value) == 1 && strings.Count(second, value) == 0 {
			result = append(result, value)
		}
	}

	for _, value := range second_arr {
		if strings.Count(second, value) == 1 && strings.Count(first, value) == 0 {
			result = append(result, value)
		}
	}

	return result
}