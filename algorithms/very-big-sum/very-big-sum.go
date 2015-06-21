package very_big_sum

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Scanf("%d", &numbers[i])
	}

	fmt.Println(sum(numbers))
}

func sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
