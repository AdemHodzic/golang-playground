package day11

func CalculateDeposit(deposit float32, rate int, max float32) int {

	counter := 0
	for deposit < max {
		perc := deposit / 100
		deposit += perc * float32(rate)
		counter++
	}

	return counter
}
