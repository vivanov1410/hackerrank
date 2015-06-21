package multiples_of_3_and_5

import "testing"

func TestSumOfMultiplesOf3And51(t *testing.T) {
	actual := SumOfMultiplesOf3And5(10)
	expected := 23
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSumOfMultiplesOf3And52(t *testing.T) {
	actual := SumOfMultiplesOf3And5(100)
	expected := 2318
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
