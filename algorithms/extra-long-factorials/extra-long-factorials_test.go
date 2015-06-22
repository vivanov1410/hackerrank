package extra_long_factorials

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	actual := factorial(big.NewInt(25))
	expected := new(big.Int)
	fmt.Sscan("15511210043330985984000000", expected)
	if actual.Cmp(expected) != 0 {
		t.Error("Expected", expected, "got", actual)
	}
}
