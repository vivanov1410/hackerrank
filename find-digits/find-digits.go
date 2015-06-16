package find_digits

import (
	"fmt"
	"strconv"
)

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	numbers := make([]string, numberOfTestCases)
	for i := 0; i < numberOfTestCases; i++ {
		fmt.Scan(&numbers[i])
	}

	for _, number := range numbers {
		fmt.Println(FindDigits(number))
	}
}

func FindDigits(numberString string) int {
	numberOfDigits := 0
	number, _ := strconv.Atoi(numberString)

	for _, r := range numberString {
		i, _ := strconv.Atoi(string(r))
		if i != 0 && number%i == 0 {
			numberOfDigits++
		}
	}

	return numberOfDigits
}
