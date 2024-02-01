package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alex":  29,
		"Winny": 23,
		"Maxs":  24,
		"Troll": 28,
	}

	delete(ages, "Troll")

	age, exists := ages["Troll"]
	if !exists {
		fmt.Println("Troll is not on the list")
	} else {
		fmt.Println("Troll is %d years old", age)
	}
}
