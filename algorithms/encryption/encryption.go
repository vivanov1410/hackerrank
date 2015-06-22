package encryption

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(encode(s))
}

func encode(s string) string {
	r := math.Sqrt(float64(len(s)))
	// minBoundary := math.Floor(r)
	maxBoundary := math.Ceil(r)

	rows := int(maxBoundary)
	columns := int(maxBoundary)

	arr := make([]rune, rows*columns)
	// fmt.Println(rows, columns, len(s))
	for i := 0; i < len(arr); i++ {
		if i < len(s) {
			arr[i] = rune(s[i])
		} else {
			arr[i] = ' '
		}
	}

	output := ""
	for i := 0; i < columns; i++ {
		line := ""
		for j := 0; j < rows; j++ {
			line += string(arr[j*columns+i])
		}

		output += strings.Trim(line, " ") + " "
	}

	return output[:len(output)-1]
}
