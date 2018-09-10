package main

import (
  "strconv"
  "strings"
)

func MaxSwap(number int) int {
  numberString := strconv.Itoa(number)
  arr := strings.Split(numberString, "")

  for i := 1; i < len(arr); i++ {
    current, _ := strconv.Atoi(arr[i])
    prev, _ := strconv.Atoi(arr[i - 1])
    if current > prev {
      arr[i], arr[i - 1] = arr[i - 1], arr[i]
      break
    }
  }

  result, _ := strconv.Atoi(strings.Join(arr, ""))
  return result
}
