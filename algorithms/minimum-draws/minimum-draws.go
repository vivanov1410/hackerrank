package minimum_draw

import "fmt"

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	testCases := make([]int, numberOfTestCases)
	for i := range testCases {
		fmt.Scanf("%d", &testCases[i])
	}

	for _, testCase := range testCases {
		fmt.Println(calcHowManyToRemove(testCase))
	}
}

func calcHowManyToRemove(n int) int {
	return n + 1
}
