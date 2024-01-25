package main

import "fmt"

func main() {
	defer handlerMessage()

	messages := []string{
		"messages 1",
		"messages 2",
		"messages 3",
		"messages 4",
	}

	fmt.Println(messages)
	messages[4] = "message 5"
}

func handlerMessage() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("handlerPanic() completed successfully")
}
