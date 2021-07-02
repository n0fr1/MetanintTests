package main

import (
	"fmt"
	"os"
)

type allPersons []person

type person struct {
	name   string
	age    int32
	weight float64
}

func main() {

	allPersons := GetPersons()

	file, err := os.Create("persons.dat")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	for _, value := range allPersons {
		fmt.Fprintf(file,
			"%-10s %-10d %-10.2f\n", //форматирование
			//"%s %d %f\n",
			value.name, value.age, value.weight)

	}

}

func GetPersons() allPersons {

	var people allPersons

	inputTable := []struct {
		name   string
		age    int32
		weight float64
	}{
		{name: "Kevin",
			age:    24,
			weight: 68.5,
		},
		{name: "Anna",
			age:    35,
			weight: 70.2,
		},
		{name: "Michael",
			age:    45,
			weight: 90.5,
		},
		{name: "Marina",
			age:    32,
			weight: 60.5,
		},
	}

	for _, value := range inputTable {
		if value.age <= 40 {
			person := fillPerson(value.name, value.age, value.weight) //надуманный пример.
			people = append(people, *person)                          //допустим, нужно получить слайс с определенными сотрудниками и его потом вывести в файл
		}
	}

	return people
}

func fillPerson(name string, age int32, weight float64) *person {

	return &person{
		name:   name,
		age:    age,
		weight: weight,
	}
}
