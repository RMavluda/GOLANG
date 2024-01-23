package main

import "fmt"

func main() {
	defer prtMessage()

	fmt.Println("main()")
	fmt.Println("main() 2")
	fmt.Println("main() 3")
}

func prtMessage() {
	fmt.Println("printMessage()")
}
