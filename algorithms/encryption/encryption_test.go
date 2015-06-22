package encryption

import "testing"

func TestEncode1(t *testing.T) {
	actual := encode("haveaniceday")
	expected := "hae and via ecy"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestEncode2(t *testing.T) {
	actual := encode("feedthedog")
	expected := "fto ehg ee dd"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestEncode3(t *testing.T) {
	actual := encode("chillout")
	expected := "clu hlt io"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestEncode4(t *testing.T) {
	actual := encode("ifmanwasmeanttostayonthegroundgodwouldhavegivenusroots")
	expected := "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau"
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
