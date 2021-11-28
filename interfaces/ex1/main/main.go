package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct{}

type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")

}

func main() {
	var tesla Vehicle = Car{}
	var boing Vehicle = Aircraft{}

	tesla.move()
	boing.move()
}
