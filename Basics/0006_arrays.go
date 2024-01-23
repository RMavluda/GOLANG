package main

import (
	"errors"
	"fmt"
)

func main() {
	messages := [3]string{"1", "2", "3"}

	printMessages(messages)

	fmt.Println(messages)
}

func printMessages(messages [3]string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}

	messages[1] = "5"

	fmt.Println(messages)

	return nil
}
