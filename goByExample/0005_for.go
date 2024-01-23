package main

import "fmt"

func main() {
	//base
	i := 0
	for i < 5 {
		i = i + 1
		fmt.Println(i)
	}
	//initial, condition, after
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
