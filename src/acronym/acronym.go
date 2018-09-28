package main

import (
  "fmt"
  "strings"
)

func main() {

  text := "helo i am razan"
  words := strings.Fields(text)
  for i, word := range words {
    fmt.Println(i, " => ", word[:1])
  }
} 