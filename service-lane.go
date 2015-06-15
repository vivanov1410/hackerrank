package main

import (
  "fmt"
)

var widths []int

func main() {
  var n, t int
  fmt.Scanf("%d %d", &n, &t)

  widths = make([]int, n)

  for i := 0; i < n; i++ {
    var width int
    fmt.Scanf("%d", &width)
    widths[i] = width
  }

  tests := make([][]int, t)

  for i := 0; i < t; i++ {
    var a, b int
    fmt.Scanf("%d %d", &a, &b)
    tests[i] = make([]int, 2)
    tests[i][0] = a
    tests[i][1] = b
  }

  for i := 0; i < t; i++ {
    vehicleType := calcVehicleType(tests[i][0], tests[i][1])
    fmt.Println(vehicleType)
  }
}

func calcVehicleType(a, b int) int {
  min := widths[a]
  for i := a+1; i <= b; i++ {
    if widths[i] < min {
      min = widths[i]
    }
  }

  return min
}
