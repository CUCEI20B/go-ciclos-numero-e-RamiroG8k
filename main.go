package main

import (
	"fmt"
)

func main() {
	var e, fact float64
	var n int

	e = 0
	fact = 1

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fact = fact * float64(i)
		e += (1 / fact)
	}
	e++

	fmt.Println(e)
}
