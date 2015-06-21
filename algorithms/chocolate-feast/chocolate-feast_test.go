package chocolate_feast

import (
	"fmt"
	"os"
	"testing"
)

func TestCalculateEaten1(t *testing.T) {
	actual := calculateEaten(10, 2, 5)
	expected := 6
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCalculateEaten2(t *testing.T) {
	actual := calculateEaten(12, 4, 4)
	expected := 3
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCalculateEaten3(t *testing.T) {
	actual := calculateEaten(6, 2, 2)
	expected := 5
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCalculateEaten4(t *testing.T) {
	inputFile, _ := os.Open("test4_input.txt")
	outputFile, _ := os.Open("test4_output.txt")

	var totalTestCases int
	_, _ = fmt.Fscanf(inputFile, "%d", &totalTestCases)

	testCases := make([]TestCase, totalTestCases)
	output := make([]int, totalTestCases)
	for i := range testCases {
		_, _ = fmt.Fscanf(inputFile, "%d %d %d", &testCases[i].money, &testCases[i].price, &testCases[i].discount)
		_, _ = fmt.Fscanf(outputFile, "%d", &output[i])
	}

	for i, testCase := range testCases {
		actual := calculateEaten(testCase.money, testCase.price, testCase.discount)
		expected := output[i]
		if actual != expected {
			t.Error("Expected", expected, "got", actual)
		}
	}
}

func TestCalculateEaten5(t *testing.T) {
	actual := calculateEaten(41894, 36, 640)
	expected := 1164
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
