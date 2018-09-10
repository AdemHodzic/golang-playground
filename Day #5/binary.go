package main

import (
  "strconv"
  "strings"
)

func Binary(number int64) bool {

  binary := strconv.FormatInt(number, 2)

  for i := 1;i < len(binary); i++ {
    if  strings.EqualFold(string(binary[i]), string(binary[i-1])) {
      return false
    }
  }

  return true

}
