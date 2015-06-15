package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var k int
		fmt.Scan(&k)

		fmt.Println(isCanceled(timestamps, k))
	}
}

func isCanceled(timestamps []int, k int) string {
	present := 0
	away := 0

	for _, timestamp := range timestamps {
		if timestamp > 0 {
			away++
		} else {
			present++
		}

		if present >= k {
			return "NO"
		}
	}

	return "YES"
}
