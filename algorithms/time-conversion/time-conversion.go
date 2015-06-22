package time_conversion

import (
	"fmt"
	"strconv"
)

func main() {
	var h, m, s, f string
	fmt.Scanf("%2v:%2v:%2v%2v", &h, &m, &s, &f)

	fmt.Println(convert(h, m, s, f))
}

func convert(h, m, s, f string) string {
	H, _ := strconv.Atoi(h)

	switch f {
	case "AM":
		if H == 12 {
			H = 0
		}
	case "PM":
		if H != 12 {
			H = H + 12
		}
	}

	return fmt.Sprintf("%02d:%v:%v", H, m, s)
}
