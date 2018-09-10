package main

import (
  "regexp"
  "strings"
)

func Banned(paragraph string, banned []string) string {

  patternString := createPattern(banned)
  used :=  make(map[string] int)

  pattern, _ := regexp.Compile(patternString)
  paragraph = pattern.ReplaceAllString(paragraph, "")
  paragraph = strings.ToLower(paragraph)
  newParagraph := strings.Split(paragraph, " ")

  for _, value := range newParagraph {
    temp := strings.ToLower(string(value))

    _, ok := used[temp]

    if ok {
      used[temp] = used[temp] + 1
    } else {
      used[temp] = 1

    }

  }


  return biggest(used)
}

func createPattern(banned []string) string {

  var builder strings.Builder
  initial := "[^A-Za-z0-9\\s]"

  builder.Write([]byte(initial))

  for _, value := range banned {
    builder.Write([]byte("|"))
    builder.Write([]byte(" "))
    builder.Write([]byte(value))

  }

  return builder.String()

}

func biggest(used map[string]int) string{

  max := 0
  var result string

  for k, v := range used {
    if v > max {
      result = k
      max = v
    }
  }


  return string(result)

}
