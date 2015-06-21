package perfect_hiring

import "testing"

func TestFindCandidateWithHighestRating(t *testing.T) {
	scores := []int{8, 6, 4, 6}
	actual := findCandidateWithHighestRating(94, 8, scores)
	expected := 1
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
