package main

import (
	"fmt"
)

func main() {
	var value int
	fmt.Scan(&value)

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	for i, v := range arr {
		if v == value {
			fmt.Println(i)
			break
		}
	}
}
