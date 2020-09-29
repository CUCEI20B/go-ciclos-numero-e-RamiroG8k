package main

import "fmt"

func main() {
	var e, fact float64
	var n int

	fmt.Scan(&n)

	fact = 1
	for i := 1; i < n; i++ {
		fact = fact * float64(i)
		e += float64(i) / fact
	}

	fmt.Println(e)
}
