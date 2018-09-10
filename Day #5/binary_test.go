package main

import (
  "testing"
  "fmt"
)

func TestBinary(t *testing.T) {

  var tests = []struct {
    input int64
    expected bool
  } {
    {5, true},
    {7, false},
    {11, false},
    {10, true},
  }

  for index ,test := range tests {
    if Binary(test.input) != test.expected {
      fmt.Printf("Failed at test #%d\n", index + 1)
      fmt.Printf("GOT %v\n", Binary(test.input))
      t.Fail()
    }
  }

}
