package sherlock_and_squares

import (
	"fmt"
	"math"
)

type testCase struct {
	a int
	b int
}

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	testCases := make([]testCase, numberOfTestCases)
	for i := 0; i < numberOfTestCases; i++ {
		fmt.Scanf("%d %d", &testCases[i].a, &testCases[i].b)
	}

	for _, testCase := range testCases {
		fmt.Println(numberOfSquareIntegers(testCase.a, testCase.b))
	}
}

func numberOfSquareIntegers(a, b int) int {
	return int(math.Floor(math.Sqrt(float64(b)))-math.Ceil(math.Sqrt(float64(a)))) + 1
}
