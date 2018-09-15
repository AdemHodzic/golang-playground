package main 

func BestRotation(arr []int ) int {

	results := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		results[i] = calcValue(arr)
		first := arr[0]
		rest := arr[1:]

		arr = arr[:0]
		arr = append(arr, rest...)
		arr = append(arr, first)
	

	}

	return findHighest(results)

}

func findHighest(results  map[int]int) int {

	final := 0
	maximum := 0

	for key, value := range results {

		if value > maximum {
			maximum = value
			final = key
		} else if value == maximum && key < final {
			final = key
		}

	}

	return final
}

func calcValue(arr []int) int {

	result := 0

	for index, value := range arr {
		if index >= value {
			result++
		}
	}

	return result

}