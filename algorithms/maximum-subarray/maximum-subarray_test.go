package maximum_subarray

import "testing"

func TestMaxSumOfContiguousSubarray1(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{1, 2, -3, 1, 1})
	expected := 3
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray2(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{1, 2, -3, -100, 100})
	expected := 100
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray3(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{1, -3, 6, 1, 2})
	expected := 9
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray4(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{-1, -3, -5, -1, -2})
	expected := -1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray5(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{1, 2, 3, 4})
	expected := 10
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray6(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{2, -1, 2, 3, 4, -5})
	expected := 10
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray7(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{1})
	expected := 1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray8(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{-1, -2, -3, -4, -5, -6})
	expected := -1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray9(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{1, -2})
	expected := 1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray10(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{1, 2, 3})
	expected := 6
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray11(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{-10})
	expected := -10
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfContiguousSubarray12(t *testing.T) {
	actual := MaxSumOfContiguousSubarray([]int{1, -1, -1, -1, -1, 5})
	expected := 5
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray1(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{1, 2, -3, 1, 1})
	expected := 5
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray2(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{1, 2, 3, 4})
	expected := 10
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray3(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{2, -1, 2, 3, 4, -5})
	expected := 11
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray4(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{1})
	expected := 1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray5(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{-1, -2, -3, -4, -5, -6})
	expected := -1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray7(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{1, -2})
	expected := 1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray8(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{1, 2, 3})
	expected := 6
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray9(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{-10})
	expected := -10
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMaxSumOfNonContiguousSubarray10(t *testing.T) {
	actual := MaxSumOfNonContiguousSubarray([]int{1, -1, -1, -1, -1, 5})
	expected := 6
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
