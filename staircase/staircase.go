package main

import "fmt"

func main() {
	var height int
	fmt.Scan(&height)

	DrawStaircase(height)
}

func DrawStaircase(height int) {
	numberOfSpaces := height - 1
	numberOfStairs := 1

	for i := 0; i < height; i++ {
		for j := 0; j < numberOfSpaces; j++ {
			fmt.Printf(" ")
		}

		for j := 0; j < numberOfStairs; j++ {
			fmt.Printf("#")
		}

		fmt.Printf("\n")
		numberOfStairs++
		numberOfSpaces--
	}
}
