package main

import "fmt"

func main() {
	messages := make([]string, 100)
	fmt.Println(len(messages))
	fmt.Println(cap(messages))
	messages = append(messages, "101")
	fmt.Println(len(messages))
	fmt.Println(cap(messages))
}
