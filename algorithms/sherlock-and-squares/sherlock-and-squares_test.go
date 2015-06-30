package sherlock_and_squares

import "testing"

func TestNumberOfSquareIntegers1(t *testing.T) {
	actual := numberOfSquareIntegers(3, 9)
	expected := 2
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestNumberOfSquareIntegers2(t *testing.T) {
	actual := numberOfSquareIntegers(17, 24)
	expected := 0
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestNumberOfSquareIntegers3(t *testing.T) {
	actual := numberOfSquareIntegers(1, 1000000000)
	expected := 31622
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
