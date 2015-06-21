package multiples_of_3_and_5

import "fmt"

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	testCases := make([]int, numberOfTestCases)
	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, testCase := range testCases {
		fmt.Println(SumOfMultiplesOf3And5(testCase))
	}
}

func SumOfMultiplesOf3And5(n int) int {
	a := 3
	b := 5

	return sumOfMiltiplesOf(n, a) + sumOfMiltiplesOf(n, b) - sumOfMiltiplesOf(n, a*b)
}

func sumOfMiltiplesOf(max, number int) int {
	n := max / number
	if max%number == 0 {
		n -= 1
	}
	a1 := number
	an := n * number
	sum := (a1 + an) * n / 2

	return sum
}
