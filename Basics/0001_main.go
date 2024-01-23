package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3

	a, _, c = 1, 6, 7

	fmt.Println(a, b, c)

}
