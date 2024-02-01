package main

import (
	"fmt"
)

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("The car is moving")
}

func (a Aircraft) move() {
	fmt.Println("The plane is flying")
}
func main() {
	tesla := Car{}
	boing := Aircraft{}
	drive(tesla)
	drive(boing)
}
