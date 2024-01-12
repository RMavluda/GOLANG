package main

import "fmt"

func main() {
	fmt.Println(findMax(1, 2, 4, -567, 394, 95860, -4, 83))
}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return max
}
