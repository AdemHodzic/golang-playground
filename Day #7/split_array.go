package main

func Split(arr []int) bool {


	var first_arr []int
	var second_arr []int
	

	var first_sum int
	var second_sum int


	for _, value := range arr {

		if first_sum <= second_sum{
			first_arr = append(first_arr, value)
			first_sum += value
		} else {
			second_arr = append(first_arr, value)
			second_sum += value
		}

	}

	return (first_sum / len(first_arr)) == (second_sum / len(second_arr))

}