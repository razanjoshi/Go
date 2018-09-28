package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

func main() {
  
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  text, _ := reader.ReadString('\n')
  words := strings.Fields(text)
  for i, word := range words {
    fmt.Println(i, " => ", word[:1])
  }
} 


// Output:
// Enter text:
// Ruby on Rails
// 0 => R
// 1 => o
// 2 => R
