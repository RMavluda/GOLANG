package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasya":  15,
		"Petya":  23,
		"Kostya": 48,
	}

	age, exists := users["Petya"]
	fmt.Println(age, exists)

	users["Sergey"] = 21

	//var users map[string]int
	//users = make(map[string]int)
	//users["Vasya"] = 19

	delete(users, "Vasya")

	for key, value := range users {
		fmt.Println(key, value)
	}

	//fmt.Println(cap(users)) -> ERROR: .\0012_maps.go:27:18: invalid argument: users (variable of type map[string]int) for cap

}
