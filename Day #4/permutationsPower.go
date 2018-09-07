package main

import (
	"math"
	"strings"
	"sort"
	"strconv"
)


func PermutationsPower(number int ) bool{

	numberStr := strconv.Itoa(number)
	if strings.Count(numberStr, "0") > 0 {
		return false
	}

	maxNumber := math.Pow(10, float64(len(numberStr)))

	var arr []int

	populate(&arr, int(maxNumber))

	numberStrArr := strings.Split(numberStr,"")
	sort.Strings(numberStrArr) 


	for _, value := range arr {
		temp := strconv.Itoa(value)
		tempArr := strings.Split(temp,"")
		sort.Strings(tempArr)


		if equal(tempArr, numberStrArr) {
			return true
		}
	}


	return false
}


func populate(arr *[]int, maxNumber int) {
	for i := 0; i < maxNumber; i++ {
		*arr = append(*arr, int(math.Pow(2, float64(i))))
	}
}

func equal(first []string, second []string) bool{

	if len(first) != len(second) {
		return false
	}

	for i, value := range first {
		if value != second[i] {
			return false
		}
	}

	return true

}