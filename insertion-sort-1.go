package insertion_sort_1

import (
	"fmt"
	"strings"
)

func main() {
	var size int
	fmt.Scan(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	InsertionSort(arr)
}

func InsertionSort(arr []int) []int {
	v := arr[len(arr)-1]

	if len(arr) == 2 {
		arr[1] = arr[0]
		arr[0] = v
	} else {
		for i := len(arr) - 1; i >= 1; i-- {
			next := arr[i-1]

			if v < next {
				arr[i] = arr[i-1]
				printPretty(arr)
			} else {
				arr[i] = v
				break
			}
		}
	}

	printPretty(arr)

	return arr
}

func printPretty(arr []int) {
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}
