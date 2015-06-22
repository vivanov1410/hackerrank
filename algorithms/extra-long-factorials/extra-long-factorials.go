package extra_long_factorials

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	fmt.Scan(&n)
	fmt.Println(factorial(big.NewInt(n)))
}

func factorial(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

	return new(big.Int).Mul(n, factorial(new(big.Int).Sub(n, big.NewInt(1))))
}
