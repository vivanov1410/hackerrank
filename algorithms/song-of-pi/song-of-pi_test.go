package song_of_pi

import "testing"

func TestLargestDecentNumber1(t *testing.T) {
	actual := isPieSong("Can I have a large container of coffee right now")
	expected := IS_PIE_SONG
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestLargestDecentNumber2(t *testing.T) {
	actual := isPieSong("Can I have a large container of tea right now")
	expected := IS_NOT_PIE_SONG
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestLargestDecentNumber3(t *testing.T) {
	actual := isPieSong("Now I wish I could recollect pi Eureka cried the great inventor Christmas Pudding Christmas Pie Is the problems very center")
	expected := IS_PIE_SONG
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestLargestDecentNumber4(t *testing.T) {
	actual := isPieSong("How I wish I could recollect pi easily today")
	expected := IS_PIE_SONG
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
