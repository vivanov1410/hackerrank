package fraud_prevention

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(encode(s))
}
