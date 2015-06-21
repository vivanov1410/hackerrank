package find_digits

import "testing"

func TestFindDigits1(t *testing.T) {
	actual := FindDigits("12")
	expected := 2
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestFindDigits2(t *testing.T) {
	actual := FindDigits("1012")
	expected := 3
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
