package maximum_subarray

import "testing"

func TestDrawStaircase1(t *testing.T) {
	actual := DrawStaircase(1)
	expected := 3
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
