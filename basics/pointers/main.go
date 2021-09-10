package main

import "fmt"

func main() {
	exampleNotPointer()
	examplePointer()
}

//a - остается в итоге прежним. Потому что передаем копию значения
func exampleNotPointer() {

	var a = 8
	fmt.Println("a before: ", a)
	incrementNotPointer(a)

	fmt.Println("a after: ", a)
}

func incrementNotPointer(x int) {
	fmt.Println("x before: ", x)
	x += 20
	fmt.Println("x after: ", x)
}

//a - в итоге меняется, потому что передаем указатель
func examplePointer() {

	var a = 8
	fmt.Println("a before: ", a)
	incrementPointer(&a)

	fmt.Println("a after: ", a)
}

func incrementPointer(x *int) {

	fmt.Println("x before: ", *x)
	*x += 20
	fmt.Println("x after: ", *x)

}
