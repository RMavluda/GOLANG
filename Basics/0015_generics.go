package main

import "fmt"

type Number interface {
	int64 | float64
}

type Userr struct {
	email string
	name  string
}

func main() {
	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"ac", "qw", "dfg"}

	d := []Userr{
		{
			email: "asd@gmail.com",
			name:  "Nika",
		},
		{
			email: "as12d@gmail.com",
			name:  "Toni",
		},
		{
			email: "as3423d@gmail.com",
			name:  "Tyger",
		},
	}

	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(searchElement(c, "qw"))
	fmt.Println(searchElement(d, Userr{
		email: "as3423d@gmail.com",
		name:  "Tyger",
	}))

	printAny(d)
}
func sum[V Number](input []V) V {
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
