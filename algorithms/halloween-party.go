package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	cuts := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scan(&cuts[i])
	}

	for _, cut := range cuts {
		fmt.Println(countPieces(cut))
	}
}

func countPieces(cut int) int {
	reminder := cut % 2
	a := cut / 2
	if reminder == 0 {
		return a * a
	} else {
		return a*a + a
	}
}
