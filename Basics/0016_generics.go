package main

import "fmt"

type Number interface {
	int64 | float64
}

type Usrr struct {
	email string
	name  string
}

func main() {

	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	d := []Usrr{
		{
			email: "abz@gmail.com",
			name:  "Nika",
		},
		{
			email: "abc345@gmail.com",
			name:  "Tamara",
		},
		{
			email: "pari@gmail.com",
			name:  "Mali",
		},
	}

	fmt.Println(compareTo(a))
	fmt.Println(compareTo(b))
	fmt.Println(searchElement(c, "2"))
	fmt.Println(searchElement(d, Usrr{
		email: "pari@gmail.com",
		name:  "Mali",
	}))

	printAny(a)
}

func compareTo[V Number](input []V) V {
	var result V
	for _, number := range input {
		result += number
	}
	return result
}

func searchElement[C comparable](elements []C, searchEl C) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}
	return false
}

func printAny[A any](input A) {
	fmt.Println(input)
}
