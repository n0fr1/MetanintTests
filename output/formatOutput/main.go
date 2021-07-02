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

	person := fillPerson("Kevin", 24, 68.5) //1-й способ заполнения

	// tom := person{ //2-й способ заполнения
	// 	name:   "Tom",
	// 	age:    24,
	// 	weight: 68.5,
	// }

	file, err := os.Create("person.dat")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Fprintf(file,
		"%-10s %-10d %-10.2f\n", //форматирование
		//"%s %d %f\n",
		person.name, person.age, person.weight)

}

func fillPerson(name string, age int32, weight float64) *person {
	return &person{
		name:   name,
		age:    age,
		weight: weight,
	}
}
