package strange_grid

import "testing"

func TestFindNumber1(t *testing.T) {
	actual := FindNumber(1, 1)
	expected := 0
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestFindNumber2(t *testing.T) {
	actual := FindNumber(1, 3)
	expected := 4
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestFindNumber3(t *testing.T) {
	actual := FindNumber(3, 4)
	expected := 16
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestFindNumber4(t *testing.T) {
	actual := FindNumber(4, 2)
	expected := 13
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestFindNumber5(t *testing.T) {
	actual := FindNumber(6, 3)
	expected := 25
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
