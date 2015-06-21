package even_fibonacci_numbers

import "testing"

func TestSumOfEvenFibonacciNumbers1(t *testing.T) {
	actual := sumOfEvenFibonacciNumbers(10)
	expected := 10
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSumOfEvenFibonacciNumbers2(t *testing.T) {
	actual := sumOfEvenFibonacciNumbers(100)
	expected := 44
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
