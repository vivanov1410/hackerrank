package maximum_subarray

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	for i := 0; i < numberOfTestCases; i++ {
		var size int
		fmt.Scan(&size)

		arr := make([]int, size)
		for i := 0; i < size; i++ {
			fmt.Scanf("%d", &arr[i])
		}

		fmt.Printf("%v %v\n", MaxSumOfContiguousSubarray(arr), MaxSumOfNonContiguousSubarray(arr))
	}
}

func MaxSumOfContiguousSubarray(arr []int) int {
	currentSum := 0
	bestSum := 0
	biggestNegative := MinInt

	for _, v := range arr {
		val := currentSum + v
		if val > 0 {
			currentSum = val
		} else {
			if v > biggestNegative {
				biggestNegative = v
			}
			currentSum = 0
		}

		if currentSum > bestSum {
			bestSum = currentSum
		}
	}

	if biggestNegative != MinInt && bestSum == 0 {
		return biggestNegative
	} else {
		return bestSum
	}
}

func MaxSumOfNonContiguousSubarray(arr []int) int {
	currentSum := 0
	biggestNegative := MinInt

	for _, v := range arr {
		if v > 0 {
			currentSum += v
		} else {
			if v > biggestNegative {
				biggestNegative = v
			}
		}
	}

	if biggestNegative != MinInt && currentSum == 0 {
		return biggestNegative
	} else {
		return currentSum
	}
}
