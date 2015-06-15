package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scanf("%d", &n)

  negatives := 0
  zeros := 0
  positives := 0

  for i := 0; i < n; i++ {
    var a int
    fmt.Scanf("%d", &a)
    if a < 0 {
      negatives++
    } else if a == 0 {
      zeros++
    } else {
      positives++
    }
  }

  fmt.Printf("%.3f %.3f %.3f", float64(positives)/float64(n), float64(negatives)/float64(n), float64(zeros)/float64(n))
}
