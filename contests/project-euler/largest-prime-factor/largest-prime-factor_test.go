package largest_prime_factor

import "testing"

func TestLargestPrimeFactor1(t *testing.T) {
	actual := largestPrimeFactor(10)
	expected := 5
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestLargestPrimeFactor2(t *testing.T) {
	actual := largestPrimeFactor(17)
	expected := 17
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestLargestPrimeFactor3(t *testing.T) {
	actual := largestPrimeFactor(10000001)
	expected := 909091
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
