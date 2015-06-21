package perfect_hiring

import "fmt"

func main() {
	var numberOfCandidates, patience, damage int
	fmt.Scanf("%d %d %d", &numberOfCandidates, &patience, &damage)

	scores := make([]int, numberOfCandidates)
	for i := range scores {
		fmt.Scanf("%d", &scores[i])
	}

	fmt.Println(findCandidateWithHighestRating(patience, damage, scores))
}

func findCandidateWithHighestRating(patience int, damage int, scores []int) int {
	times := len(scores)
	if patience/damage < times {
		times = patience / damage
	}

	ratings := make([]int, times)
	for i := 0; i < times; i++ {
		ratings[i] = patience * scores[i]
		patience -= damage
	}

	highestRating := 0
	highestRatingId := -1
	for i, rating := range ratings {
		if rating > highestRating {
			highestRating = rating
			highestRatingId = i + 1
		}
	}

	return highestRatingId
}
