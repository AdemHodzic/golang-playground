package main

import (
  "testing"
  "fmt"
)

func TestFindMissing(t *testing.T) {

  var tests = []struct {
      input []int
      expected int
  } {
      {[]int {1,2,0}, 3},
      {[]int {3,4,-1,1}, 2},
      {[]int {7,8,9,11,12}, 1},
  }

  for index, test := range tests {

    if FindMissing(test.input) != test.expected {
      fmt.Printf("FAILED TEST AT #%v\n", index + 1)
      fmt.Printf("GOT: %v\n", FindMissing(test.input))
      t.Fail()
    }

  }

}
