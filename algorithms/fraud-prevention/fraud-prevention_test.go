package fraud_prevention

import "testing"

func TestEncode1(t *testing.T) {
	actual := encode("haveaniceday")
	expected := "hae and via ecy"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
