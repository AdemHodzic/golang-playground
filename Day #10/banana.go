package main 

func BananaSpeed(piles []int, hours int) int {

	maximum := max(piles)

	if len(piles) == hours {
		return maximum
	} 

	for i := min(piles);; i++ {
		if check(i, piles, hours) == true {
			return i
		}
	}

	return -1
}


func max(arr []int) int {
	maximum := 0

	for _, v := range arr {
		if v > maximum {
			maximum = v
		}
	}

	return maximum
}

func min(arr []int) int {
	minimum := 99999

	for _, v := range arr {
		if v < minimum {
			minimum = v
		}
	}

	return minimum

}

func check(speed int, piles []int, hours int) bool {

	index := 0

	for i := 0; i < hours; i++ {

		piles[index] -=  speed
		
		if piles[index] <= 0 {
			index++
		}

		if index == len(piles) {
			return true
		}
	}

	return false

}