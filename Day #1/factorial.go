package main

func Factorialize(number int) int{
	if number == 0 {
		return 1
	}

	return number * Factorialize(number - 1)
}