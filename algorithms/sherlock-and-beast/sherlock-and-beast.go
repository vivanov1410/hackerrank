package sherlock_and_beast

import "fmt"

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	testCases := make([]int, numberOfTestCases)
	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, testCase := range testCases {
		fmt.Println(largestDecentNumber(testCase))
	}
}

func largestDecentNumber(n int) string {
	if n > 2 {
		if n%3 == 0 {
			runes := make([]rune, n)
			for i := range runes {
				runes[i] = '5'
			}
			return string(runes)
		}

		for i := 0; i <= n; i++ {
			n5 := n - i
			n3 := n - n5

			if n5%3 == 0 && n3%5 == 0 {
				runes := make([]rune, n)
				for i := 0; i < n5; i++ {
					runes[i] = '5'
				}
				for i := n5; i < n; i++ {
					runes[i] = '3'
				}
				return string(runes)
			}
		}
	}

	return "-1"
}
