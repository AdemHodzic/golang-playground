package main

import (
  "sort"
)

func FindMissing(arr []int) int {

  sort.Ints(arr)

  counter := 1

  for 1 > 0 {
    if arrHasCounter(arr, counter) == false {
      return counter
    } else {
      counter++
    }
  }

  return -1
}

func arrHasCounter(arr []int, number int) bool {

  for _, value := range arr{
    if value == number {
      return true
    }
  }

  return false

}
