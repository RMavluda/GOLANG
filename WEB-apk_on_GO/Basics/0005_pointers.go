package main

import "fmt"

func main() {
	number := 5
	var p *int

	p = &number
	//message := "Skoro ya stanu nidzya!" //0xc000010230

	fmt.Println(p)
	fmt.Println(&number)

	*p = 10

	fmt.Println(number)

	//changemessage(message)

	//fmt.Println(&message)
}

func changemessage(message string) { //0xc000010230
	// *message += " ( in loop changeMessage())"
	fmt.Println(&message)
}
