package main

import (
	"fmt"
	"math/rand"
)

func SelectionSortRecAux(a []int, n int) []int {
	if n == 0 || n == 1 {
		return a
	} else {
		max := 0
		for i := 0; i < n; i++ {
			if a[i] > a[max] {
				max = i
			}
		}
		tmp := a[n-1]
		a[n-1] = a[max]
		a[max] = tmp
	}
	return SelectionSortRecAux(a, n-1)
}

func SelectionSortRec(a []int) []int {
	return SelectionSortRecAux(a, len(a))
}

func SelectionSortIter(a []int) []int {
	if len(a) == 0 {
		return a
	}
	var min int = 0
	var idx int = 0
	for idx != len(a)-1 {
		for i := idx; i < len(a); i++ {
			if a[i] < a[min] {
				min = i
			}
		}
		if min != idx {
			var tmp int
			tmp = a[idx]
			a[idx] = a[min]
			a[min] = tmp
		}
		idx++
		min = idx
	}
	return a
}

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(10)
	}
	fmt.Printf("a: %v\n", a)
	SelectionSortRec(a[:])
	fmt.Printf("a: %v\n", a)
}
