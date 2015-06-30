package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PI              = "31415926535897932384626433833"
	IS_PIE_SONG     = "It's a pi song."
	IS_NOT_PIE_SONG = "It's not a pi song."
)

func main() {
	var numberOfTestCases int
	fmt.Scan(&numberOfTestCases)

	testCases := make([]string, numberOfTestCases)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < numberOfTestCases; i++ {
		line, _, _ := reader.ReadLine()
		testCases[i] = string(line)
	}

	for _, testCase := range testCases {
		fmt.Println(isPieSong(testCase))
	}
}

func isPieSong(song string) string {
	words := strings.Fields(song)
	for i, word := range words {
		l, _ := strconv.Atoi(string(PI[i]))
		if len(word) != l {
			return IS_NOT_PIE_SONG
		}
	}

	return IS_PIE_SONG
}
