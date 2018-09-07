package main

import (
	"strconv"
	"strings"
)

func BinaryDistance(number int)int {
	bin := strconv.FormatInt(int64(number), 2)
	

	if strings.Count(string(bin), "1") == 1 {
		return 0
	}


	result := 0

	for i := 0; i < len(bin); i++ {

		if strings.EqualFold(string(bin[i]), "1") {

			var temp = 0

			for j := i + 1; j < len(bin); j++ {
				
				temp++
				
				if strings.EqualFold(string(bin[j]), "1") {
					break
				}
			}
			if temp > result {
				result = temp
			}

		}

	}

	return result

}