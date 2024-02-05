package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}

	s := slices.IsSorted(strs)
	fmt.Println("Sorted:", s)

	slices.Sort(strs)
	fmt.Println("Strigns:", strs)

	ints := []int{5, 9, 2}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	s = slices.IsSorted(ints)
	fmt.Println("SortedL", s)
}
