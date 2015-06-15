package main

import (
	"fmt"
)

func main() {
	var actualDay, actualMonth, actualYear int
	fmt.Scanf("%d %d %d", &actualDay, &actualMonth, &actualYear)

	var expectedDay, expectedMonth, expectedYear int
	fmt.Scanf("%d %d %d", &expectedDay, &expectedMonth, &expectedYear)

	fine := 0

	if (actualDay + (actualMonth-1)*31 + (actualYear-1)*372) <= (expectedDay + (expectedMonth-1)*31 + (expectedYear-1)*372) {
		fine = 0
	} else if actualYear == expectedYear && actualMonth == expectedMonth {
		fine = (actualDay - expectedDay) * 15
	} else if actualYear == expectedYear {
		fine = (actualMonth - expectedMonth) * 500
	} else {
		fine = 10000
	}

	fmt.Println(fine)
}
