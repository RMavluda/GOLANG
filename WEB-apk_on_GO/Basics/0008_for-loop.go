package main

import "fmt"

func main() {
	//messages := []string{
	//	"messages 1",
	//	"messages 2",
	//	"messages 3",
	//	"messages 4",
	//}

	//for i := 0; i < len(messages); i++{
	//	fmt.Println(messages[i])
	//}

	//for _, messages := range messages {
	//	fmt.Println(messages)
	//}

	counter := 0
	for {
		counter++

		if counter == 100 {
			break
		}

		fmt.Println(counter)
	}
}
