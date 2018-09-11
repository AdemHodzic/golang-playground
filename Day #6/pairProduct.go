package main

func PairProduct(arr []int) int {

  result := 0

  for i := 1; i < len(arr); i++ {
    current := arr[i - 1] * arr[i]
    if current > result {
      result = current
    }
  }

  return result

}
