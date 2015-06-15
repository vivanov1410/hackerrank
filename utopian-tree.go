package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scanf("%d", &n)

  inputs := make([]int, n)

  for i := 0; i < n; i++ {
    var a int
    fmt.Scanf("%d", &a)
    inputs[i] = a
  }

  for i := 0; i < n; i++ {
    length := calcLength(inputs[i])
    fmt.Println(length)
  }
}

func calcLength(numberOfCycles int) int {
  l := 1

  for i := 1; i <= numberOfCycles; i++ {
    if i % 2 == 0 {
      l++
    } else {
      l *= 2
    }
  }

  return l
}
