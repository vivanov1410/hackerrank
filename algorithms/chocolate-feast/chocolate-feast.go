package chocolate_feast

import "fmt"

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	testCases := make([]TestCase, numberOfTestCases)
	for i := range testCases {
		fmt.Scanf("%d %d %d", &testCases[i].money, &testCases[i].price, &testCases[i].discount)
	}

	for _, testCase := range testCases {
		fmt.Println(calculateEaten(testCase.money, testCase.price, testCase.discount))
	}
}

type TestCase struct {
	money    int
	price    int
	discount int
}

func calculateEaten(money, price, discount int) int {
	eaten := 0
	totalWrappers := 0

	eaten += money / price
	totalWrappers = eaten

	for totalWrappers >= discount {
		n := totalWrappers / discount
		eaten += n
		totalWrappers += n * (1 - discount)
	}

	return eaten
}
