package main

func LargestFour(arr [][]int) (result []int) {
	for _, nums := range arr {
		result = append(result, biggest(nums))
	}
	return result
}

func biggest(arr []int) int {
	result := 0

	for _, value := range arr {
		if value > result {
			result = value
		}
	}

	return result
}