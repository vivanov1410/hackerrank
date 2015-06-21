package begin_end

import "fmt"

func main() {
	var length int
	fmt.Scan(&length)

	var s string
	fmt.Scan(&s)

	fmt.Println(sumOfSubstrings(s))
}

func sumOfSubstrings(s string) (sum int) {
	sum = 0
	chars := make(map[string]int)

	for _, char := range s {
		_, ok := chars[string(char)]
		if ok {
			chars[string(char)]++
		} else {
			chars[string(char)] = 1
		}
	}

	for _, v := range chars {
		sum += calcSum(v)
	}

	return sum
}

func calcSum(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 3
	}

	return n * (1 + n) / 2
}
