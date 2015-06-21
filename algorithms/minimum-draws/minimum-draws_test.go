package minimum_draw

import "testing"

func TestCalcHowManyToRemove1(t *testing.T) {
	actual := calcHowManyToRemove(1)
	expected := 2
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCalcHowManyToRemove2(t *testing.T) {
	actual := calcHowManyToRemove(2)
	expected := 3
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
