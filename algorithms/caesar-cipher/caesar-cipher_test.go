package caesar_cipher

import "testing"

func TestCipher1(t *testing.T) {
	actual := cipher("middle-Outz", 2)
	expected := "okffng-Qwvb"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCipher2(t *testing.T) {
	actual := cipher("www.xy.abc", 87)
	expected := "fff.gh.jkl"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestIsLowercase1(t *testing.T) {
	actual := isLowercase('a')
	expected := true
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestIsLowercas2(t *testing.T) {
	actual := isLowercase('A')
	expected := false
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestIsUppercase1(t *testing.T) {
	actual := isUppercase('a')
	expected := false
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestIsUppercas2(t *testing.T) {
	actual := isUppercase('A')
	expected := true
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
