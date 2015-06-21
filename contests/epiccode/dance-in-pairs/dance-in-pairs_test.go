package dance_in_pairs

import "testing"

func TestMaxNumberOfPairs1(t *testing.T) {
	boys := []int{1, 2, 3}
	girls := []int{12, 11, 101}
	condition := 10
	actual := maxNumberOfPairs(boys, girls, condition)
	expected := 2
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxNumberOfPairs2(t *testing.T) {
	boys := []int{12, 12, 13, 14}
	girls := []int{2, 2, 1, 5}
	condition := 10
	actual := maxNumberOfPairs(boys, girls, condition)
	expected := 3
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
