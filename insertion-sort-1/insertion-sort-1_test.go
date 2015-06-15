package insertion_sort_1

import "testing"

func TestInsertionSort1(t *testing.T) {
	actual := InsertionSort([]int{1, 2, 4, 3})
	expected := []int{1, 2, 3, 4}
	if !isEqualSlices(actual, expected) {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestInsertionSort2(t *testing.T) {
	actual := InsertionSort([]int{1})
	expected := []int{1}
	if !isEqualSlices(actual, expected) {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestInsertionSort3(t *testing.T) {
	actual := InsertionSort([]int{1, 2, 3, 5, 3})
	expected := []int{1, 2, 3, 3, 5}
	if !isEqualSlices(actual, expected) {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestInsertionSort4(t *testing.T) {
	actual := InsertionSort([]int{4, 5, 6, 10, 2})
	expected := []int{2, 4, 5, 6, 10}
	if !isEqualSlices(actual, expected) {
		t.Error("Expected", expected, "got", actual)
	}
}

func isEqualSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
