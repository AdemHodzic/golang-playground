package main

func Warmer(temperatures []int) []int {

	var result []int

	for i, value := range temperatures {

		countner := 1

		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > value {
				break
			}
			
			countner++

			if j == len(temperatures) - 1 {
				countner = 0
			}

		}

		if i == len(temperatures) - 1 {
				countner = 0
		}

		result = append(result, countner)

	}

	return result
}