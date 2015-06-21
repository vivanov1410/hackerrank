package even_fibonacci_numbers

import "fmt"

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	testCases := make([]int, numberOfTestCases)
	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, testCase := range testCases {
		fmt.Println(sumOfEvenFibonacciNumbers(testCase))
	}
}

func sumOfEvenFibonacciNumbers(number int) int {
	a := 2
	b := 3
	c := 5
	sum := 0

	for a < number {
		sum += a
		a = b + c
		b = a + c
		c = a + b
	}

	return sum
}
