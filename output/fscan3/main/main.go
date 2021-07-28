package main

import (
	"fmt"
	"os"
)

func main() {

	var name string
	var age uint32
	var height uint32

	fmt.Print("Введите имя: ")
	fmt.Scan(&name) //сразу считываем со стандартного ввода

	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age) //с указанием ридера

	fmt.Print("Введите рост: ")
	fmt.Fscanf(os.Stdin, "%d", &height) //с указанием формата строки. В обоих случаях можно указывать несколько значений.

	fmt.Printf("%s %d %d\n", name, age, height)
}
