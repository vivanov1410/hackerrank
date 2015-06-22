package time_conversion

import "testing"

func TestCipher1(t *testing.T) {
	actual := convert("07", "05", "45", "PM")
	expected := "19:05:45"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCipher2(t *testing.T) {
	actual := convert("12", "40", "22", "AM")
	expected := "00:40:22"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCipher3(t *testing.T) {
	actual := convert("11", "33", "11", "PM")
	expected := "23:33:11"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCipher4(t *testing.T) {
	actual := convert("06", "40", "03", "AM")
	expected := "06:40:03"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCipher5(t *testing.T) {
	actual := convert("12", "45", "54", "PM")
	expected := "12:45:54"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
