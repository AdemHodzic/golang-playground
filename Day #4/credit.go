package main

import (
	"strconv"
)



func CheckCard(number int ) string{

	cardNames := make(map[int] string)

	cardNames[34] = "American Express"
	cardNames[37] = "American Express"

	cardNames[51] = "MasterCard"
	cardNames[52] = "MasterCard"
	cardNames[53] = "MasterCard"
	cardNames[54] = "MasterCard"
	cardNames[55] = "MasterCard"
	

	if !isValid(number) {
		return "Invalid"
	}

	firstTwo := string((strconv.Itoa(number))[:2])

	if int(firstTwo[0]) == 4 {
		return "Vise"
	}

	for k,v := range  cardNames {
		num, err := strconv.Atoi(firstTwo)
		
		if err != nil {
			panic("Fatal error")
		}

		if k == num {
			return v
		}
	}
	return "Invalid"
}

func isValid(number int) bool {

	first := 0
	second := 0

	str := strconv.Itoa(number)

	for index, value := range str {

		if index % 2 == 0 {
			second += int(value)
		} else {
			first += int(value)
		}

	}
	return (first + second) % 10 == 0
}