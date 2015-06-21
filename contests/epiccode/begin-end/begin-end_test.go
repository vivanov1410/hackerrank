package begin_end

import "testing"

func TestSumOfSubstrings1(t *testing.T) {
	actual := sumOfSubstrings("a")
	expected := 1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSumOfSubstrings2(t *testing.T) {
	actual := sumOfSubstrings("ab")
	expected := 2
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSumOfSubstrings3(t *testing.T) {
	actual := sumOfSubstrings("aba")
	expected := 4
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSumOfSubstrings4(t *testing.T) {
	actual := sumOfSubstrings("abab")
	expected := 6
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSumOfSubstrings5(t *testing.T) {
	actual := sumOfSubstrings("ababa")
	expected := 9
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSumOfSubstrings6(t *testing.T) {
	actual := sumOfSubstrings("ababac")
	expected := 10
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestSumOfSubstrings7(t *testing.T) {
	actual := sumOfSubstrings("ababaca")
	expected := 14
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
