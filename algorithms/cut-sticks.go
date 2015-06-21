package main

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

var stickLengths []int

func main() {
	var n int
	fmt.Scanf("%d", &n)

	stickLengths = make([]int, n)

	for i := 0; i < n; i++ {
		var stickLength int
		fmt.Scanf("%d", &stickLength)
		stickLengths[i] = stickLength
	}

	for min := findMinimumStick(); min != 0; min = findMinimumStick() {
		fmt.Println(cutSticks(min))
	}
}

func findMinimumStick() int {
	min := MaxInt

	for _, stickLength := range stickLengths {
		if stickLength != 0 && stickLength < min {
			min = stickLength
		}
	}

	if min == MaxInt {
		return 0
	} else {
		return min
	}
}

func cutSticks(cutLength int) int {
	numberOfCuttedSticks := 0
	l := len(stickLengths)

	for i := 0; i < l; i++ {
		if stickLengths[i] != 0 {
			stickLengths[i] -= cutLength
			numberOfCuttedSticks++
		}
	}

	return numberOfCuttedSticks
}
