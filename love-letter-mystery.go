package main

import (
	"fmt"
	"math"
)

var widths []int

func main() {
	var t int
	fmt.Scanf("%d", &t)

	tests := make([]string, t)

	for i := 0; i < t; i++ {
		var test string
		fmt.Scan(&test)
		tests[i] = test
	}

	for i := 0; i < t; i++ {
		fmt.Println(countOperations(tests[i]))
	}
}

func countOperations(s string) int {
	runes := []rune(s)
	numberOfOperations := 0

	times := len(runes) / 2
	lastIndex := len(runes) - 1

	for i := 0; i < times; i++ {
		if runes[i] != runes[lastIndex-i] {
			numberOfOperations += int(math.Abs(float64(runes[i]) - float64(runes[lastIndex-i])))
		}
	}

	return numberOfOperations
}
