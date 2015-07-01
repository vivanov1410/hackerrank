package grid_search

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestHasPattern1(t *testing.T) {
	grid := []string{
		"7283455864",
		"6731158619",
		"8988242643",
		"3830589324",
		"2229505813",
		"5633845374",
		"6473530293",
		"7053106601",
		"0834282956",
		"4607924137",
	}
	pattern := []string{
		"9505",
		"3845",
		"3530",
	}
	actual := hasPattern(grid, pattern)
	expected := "YES"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasPattern2(t *testing.T) {
	grid := []string{
		"400453592126560",
		"114213133098692",
		"474386082879648",
		"522356951189169",
		"887109450487496",
		"252802633388782",
		"502771484966748",
		"075975207693780",
		"511799789562806",
		"404007454272504",
		"549043809916080",
		"962410809534811",
		"445893523733475",
		"768705303214174",
		"650629270887160",
	}
	pattern := []string{
		"99",
		"99",
	}
	actual := hasPattern(grid, pattern)
	expected := "NO"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHasPattern3(t *testing.T) {
	inputFile, _ := os.Open("input.txt")
	outputFile, _ := os.Open("output.txt")

	var numberOfTestCases int
	reader := bufio.NewReader(inputFile)
	line, _ := reader.ReadString('\n')
	fmt.Sscanf(line, "%d", &numberOfTestCases)

	testCases := make([]string, numberOfTestCases)
	for i := 0; i < numberOfTestCases; i++ {
		var numberOfRows, numberOfColumns int

		line, _ := reader.ReadString('\n')
		fmt.Sscanf(line, "%d %d", &numberOfRows, &numberOfColumns)

		grid := readMatrixFromFile(reader, numberOfRows, numberOfColumns)

		line, _ = reader.ReadString('\n')
		fmt.Sscanf(line, "%d %d", &numberOfRows, &numberOfColumns)
		pattern := readMatrixFromFile(reader, numberOfRows, numberOfColumns)
		testCases[i] = hasPattern(grid, pattern)

		if i == 0 {
			fmt.Println(testCases[i])
			fmt.Println(grid[0], grid[len(grid)-1])
		}
	}

	for i, testCase := range testCases {
		var expected string
		fmt.Fscanln(outputFile, &expected)
		actual := testCase
		if actual != expected {
			t.Error("Expected", expected, "got", actual, "test case:", i)
		}
	}
}
