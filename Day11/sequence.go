package day11

func Sequence(arr []int) bool {

	for i := 0; i < len(arr); i++ {

		temp := build(arr, i)
		if check(temp) == true {
			return true
		}

	}

	return false
}

func check(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

func build(arr []int, index int) []int {

	temp := make([]int, len(arr)-1)
	ind := 0
	for i, v := range arr {
		if i == index {
			continue
		}
		temp[ind] = v
		ind++
	}

	return temp

}
