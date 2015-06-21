package white_falcon_and_sequence

import "testing"

func TestMaxSum1(t *testing.T) {
	numbers := []int{1, 7, 4, 0, 9, 4, 0, 1, 8, 8, 2, 4}
	// x={4,0,9,4}
	// y={8,8,2,4}
	actual := maxSum(numbers)
	expected := 120
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSum2(t *testing.T) {
	numbers := []int{1, 2, 3, 0, 5, 3}
	actual := maxSum(numbers)
	expected := 21
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSum3(t *testing.T) {
	numbers := []int{1, 1}
	actual := maxSum(numbers)
	expected := 1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSum4(t *testing.T) {
	numbers := []int{1, 1, 2}
	actual := maxSum(numbers)
	expected := 2
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSum5(t *testing.T) {
	numbers := []int{0, -4, 2}
	actual := maxSum(numbers)
	expected := 0
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSum6(t *testing.T) {
	numbers := []int{-2, 5, -3}
	actual := maxSum(numbers)
	expected := 6
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
