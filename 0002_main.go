package main

import (
	"errors"
	"fmt"
)

func main() {
	/*var message string
	message = sayHello("Max", 21)

	printMessage(message)*/
	message, err := enterTheClub(25)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
	// fmt.Println(enterTheClub(15))
}
func printMessage(message string) {
	fmt.Println(message)
}
func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s! You are %d years old!", name, age)
}
func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Welcome, just carefully", nil
	} else if age >= 45 && age < 65 {
		return "Do you really like this music?", nil
	} else if age >= 65 {
		return "This is too much for you!", errors.New("you are too old")
	}
	return "You are under 18", errors.New("you are too young")
}
