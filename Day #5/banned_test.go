package main

import (
  "testing"
  "strings"
  "fmt"
)


func TestBanned(t *testing.T) {

  var tests = []struct{
    paragraph string
    banned []string
    expected string
  } {
    {"Bob hit a ball, the hit BALL flew far after it was hit.", []string {"hit"},"ball"},
  }

  for index, test := range tests {

    if !strings.EqualFold(Banned(test.paragraph, test.banned), test.expected) {
      fmt.Printf("Failed at test #%d\n", index + 1)
      fmt.Printf("GOT %s\n", Banned(test.paragraph, test.banned))

      t.Fail()
    }

  }

}
