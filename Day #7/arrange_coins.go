package main 

func ArrangeCoins(coins int) int {

	counter := 0
	number := 1
	for coins > 0 {
		coins -= number
		if coins > 0 {
			counter++
		}
		number++
	}

	return counter

}