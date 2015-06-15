package strange_grid

import (
	"fmt"
)

func main() {
	var row, column int
	fmt.Scanf("%d %d", &row, &column)

	fmt.Println(FindNumber(row, column))
}

func FindNumber(row, column int) int {
	number := 0

	oddRow := []int{0, 0, 2, 4, 6, 8}
	evenRow := []int{0, 1, 3, 5, 7, 9}

	if isEven(row) {
		number = evenRow[column] + (row/2-1)*10
	} else {
		number = oddRow[column] + (row/2)*10
	}

	return number
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}
