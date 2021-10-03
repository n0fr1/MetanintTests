package main

import (
	"fmt"
)

func main() {
	simpleExamplePointer()
	clearPointer()
	exampleNotPointer()
	examplePointer()
	exampleNewPointer()
	returnPointerExample()
}

//простой пример с использованием указателя
func simpleExamplePointer() {
	f := 2.3
	pf := &f

	fmt.Println("Address:", pf)
	fmt.Println("Value:", *pf)
}

//Если указателю не присвоен адреc объекта, то такой указатель
//по умолчанию имеет значение nil
func clearPointer() {
	var pf *float64

	if pf != nil {
		fmt.Println("Value:", *pf) //если написать без условия - вывалится в ошибку
	} else {
		fmt.Println("Value:", "nil")
	}
}

//a - остается в итоге прежним. Потому что передаем копию значения
func exampleNotPointer() {
	var a = 8
	fmt.Println("a before: ", a)
	incrementNotPointer(a)
	fmt.Println("a after: ", a)
	incrementNotPointer(a)
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

//для понимания - используем функцию New. Объект, созданный с помощью new
//ничем не отличается от обычной переменной. Но чтобы обратиться к нему (получить, изменить адрес) - необходимо использовать указатель.
func exampleNewPointer() {
	var a = 8
	p := new(int)
	*p = a

	fmt.Println("a before: ", a)
	incrementPointer(p)
	fmt.Println("a after: ", a)
}

func incrementPointer(x *int) {
	fmt.Println("x before: ", *x)
	*x += 20
	fmt.Println("x after: ", *x)
}

//функция может возвращать указатель.
func returnPointerExample() {
	p1 := createPointer(7)
	fmt.Println("p1", *p1)

	p2 := createPointer(10)
	fmt.Println("p2", *p2)
}

func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}
