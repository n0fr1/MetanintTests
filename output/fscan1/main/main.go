package main

import (
	"fmt"
	"os"
)

type person struct {
	name   string
	age    int32
	weight float64
}

func main() {
	filename := "person.dat"
	writeData(filename)
	readData(filename)
}

func writeData(filename string) {

	tom := person{name: "Tom", age: 24, weight: 68.5}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s %d %.2f\n", tom.name, tom.age, tom.weight) //сохраняем данные в файл
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readData(filename string) {

	tom := person{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	//считывание данных из файла в структуру
	_, err = fmt.Fscanf(file, "%s %d %f\n", &tom.name, &tom.age, &tom.weight)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%-8s %-8d %-8.2f\n", tom.name, tom.age, tom.weight)
}
