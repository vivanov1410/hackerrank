package white_falcon_and_sequence

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Scanf("%d", &numbers[i])
	}

	fmt.Println(maxSum(numbers))
}

func maxSum(numbers []int) int {
	max := 0

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}

	for t := (len(numbers) - 1) / 2; t >= 1 && math.Pow(9.0, float64(t)) > float64(max); t-- {
		size := t
		xIndex := 0

		for xIndex+size*2+1 <= len(numbers) {
			yIndex := xIndex + size + 1
			// fmt.Println("xIndex", xIndex, "yIndex", yIndex)
			for yIndex+size <= len(numbers) {
				x := numbers[xIndex : xIndex+size]
				y := numbers[yIndex : yIndex+size]
				// fmt.Println("x", x)
				// fmt.Println("y", y)

				sum := calcSum(x, y)
				if sum > max {
					max = sum
					// fmt.Println("x", x)
					// fmt.Println("y", y)
					// fmt.Println("max", max)
				}

				yIndex++
			}

			xIndex++
		}
	}

	return max
}

func calcSum(x, y []int) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i] * y[len(y)-i-1]
	}

	return sum
}
