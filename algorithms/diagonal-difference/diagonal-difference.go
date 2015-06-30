package main

import (
  "fmt"
  "math"
)

func main() {
  var n int
  fmt.Scanf("%d", &n)

  leftIndex := 0
  rightIndex := n - 1
  sum := 0

  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      var a int
      fmt.Scanf("%d", &a)
      if j == leftIndex {
        // fmt.Println(a)
        sum += a
      }
      if j == rightIndex {
        // fmt.Println(a)
        sum -= a
      }
    }

    leftIndex++
    rightIndex--
  }

  fmt.Println(math.Abs(float64(sum)))
}
