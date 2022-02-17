package main

import "fmt"

type greeter interface {
	greet(string) string
}

type russian struct{}
type american struct{}

func (r *russian) greet(name string) string {
	return fmt.Sprintf("Привет, %s %s", name, "!\n")
}

func (a *american) greet(name string) string {
	return fmt.Sprintf("Hello, %s%s", name, "!\n")
}

func sayHello(g greeter, name string) {
	fmt.Printf("%s", g.greet(name))
}

func main() {
	sayHello(&russian{}, "Петя")
	sayHello(&american{}, "John")
}
