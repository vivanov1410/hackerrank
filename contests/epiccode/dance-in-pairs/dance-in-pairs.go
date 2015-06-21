package dance_in_pairs

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var numberOfPeople, condition int
	fmt.Scanf("%d %d", &numberOfPeople, &condition)

	boyHeights := make([]int, numberOfPeople)
	girlHeights := make([]int, numberOfPeople)

	for i := range boyHeights {
		fmt.Scanf("%d", &boyHeights[i])
	}
	for i := range girlHeights {
		fmt.Scanf("%d", &girlHeights[i])
	}

	fmt.Println(maxNumberOfPairs(boyHeights, girlHeights, condition))
}

func maxNumberOfPairs(boyHeights []int, girlHeights []int, condition int) int {
	pairs := 0

	sort.Sort(sort.Reverse(sort.IntSlice(boyHeights)))
	sort.Sort(sort.Reverse(sort.IntSlice(girlHeights)))

	fmt.Println(boyHeights)
	fmt.Println(girlHeights)

	boyIndex := 0
	girlIndex := 0
	for boyIndex < len(boyHeights) && girlIndex < len(girlHeights) {
		if math.Abs(float64(boyHeights[boyIndex]-girlHeights[girlIndex])) <= float64(condition) {
			pairs++
			boyIndex++
			girlIndex++
		} else {
			if boyHeights[boyIndex] > girlHeights[girlIndex] {
				boyIndex++
			} else {
				girlIndex++
			}
		}
	}

	return pairs
}
