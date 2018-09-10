package main

import (
  "testing"
  "fmt"
)

func TestMaxSwap(t *testing.T) {

  var tests = []struct{
    input int
    expected int
  } {
    {2736,7236},
    {9973,9973},
  }

  for index ,test := range tests {
    if MaxSwap(test.input) != test.expected {
      fmt.Printf("Failed at test #%d\n", index + 1)
      fmt.Printf("GOT %d\n", MaxSwap(test.input))
      t.Fail()
    }
  }

}
