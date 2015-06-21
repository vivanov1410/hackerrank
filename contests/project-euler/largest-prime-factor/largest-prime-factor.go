package largest_prime_factor

import (
	"fmt"
	"math"
)

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	testCases := make([]int, numberOfTestCases)
	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, testCase := range testCases {
		fmt.Println(largestPrimeFactor(testCase))
	}
}

func largestPrimeFactor(limit int) int {
	n := limit

	lastFactor := 1
	if n%2 == 0 {
		lastFactor = 2
		n = n / 2
		for n%2 == 0 {
			n = n / 2
		}
	}

	factor := 3
	for n > 1 || float64(factor) <= math.Sqrt(float64(n)) {
		if n%factor == 0 {
			lastFactor = factor
			n = n / factor
			for n%factor == 0 {
				n = n / factor
			}
		}
		factor += 2
	}

	return lastFactor
}
