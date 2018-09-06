package main

import "sort"

func Inserter(arr []int, number int) int {
	
	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)

	if arr[0] >= number {
		return 0
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] >= number && arr[i-1] < number {
			return i
		}
	}

	return len(arr)
}