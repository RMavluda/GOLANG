package main

import (
	"fmt"
	"time"
)

func main() {

	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "Kanal 1. Proshlo 200 ms"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			message2 <- "Kanal 2. Proshlo 1 s"
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		default:
		}
	}
	/*msg := make(chan string, 3) //buferizirovanniy  kanal

	msg <- "Kanal Ninja"
	msg <- "Kanal Ninja 2"
	msg <- "Kanal Ninja 3"

	close(msg)
	//for {
	//	value, ok := <-msg
	//	if !ok {
	//		break
	//	}
	//
	//	fmt.Println(value)
	//}

	for m := range msg {
		fmt.Println(m)
	}*/
}
