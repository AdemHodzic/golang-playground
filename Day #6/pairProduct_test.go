package main

import (
  "testing"
  "fmt"
)

func TestPairProduct(t *testing.T) {

  var tests = []struct {
    input []int
    expected int
  } {
    {[]int {2,3,-2,4}, 6},
    {[]int {-2,0,-1}, 0},
  }

  for index, test := range tests {

    if PairProduct(test.input) != test.expected {
      fmt.Printf("FAILED TEST AT #%v\n", index + 1)
      fmt.Printf("GOT: %v\n", PairProduct(test.input))
      t.Fail()
    }

  }

}
