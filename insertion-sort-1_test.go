package insertion_sort_1

import "testing"

func TestInsertionSort1(t *testing.T) {
	actual := InsertionSort([]int{1, 2, 4, 3})
	expected := []int{1, 2, 3, 4}
	if actual != expected {
		t.Error("Wrong", actual)
	}
}
