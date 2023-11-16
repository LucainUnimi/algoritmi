package main

import (
	"fmt"
)

func insertOrd(a []int, n int) []int {
	if len(a) == 0 {
		return append(a, n)
	}
	for i := range a {
		if n < a[i] {
			return append(a[:i], append([]int{n}, a[i:]...)...)
		}
	}
	return append(a, n)
}

func main() {
	var numbers []int
	var number int
	for {
		fmt.Scan(&number)
		if number == 0 {
			break
		}
		numbers = insertOrd(numbers, number)
	}

	fmt.Printf("numbers: %v\n", numbers)
}
