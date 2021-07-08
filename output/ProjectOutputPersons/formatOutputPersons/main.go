package main

import (
	"fmt"
	"os"

	"github.com/n0fr1/MetanintTests/persons"
)

func main() {

	// inputTable := []struct { //альтернативный вариант создания и заполнения структуры
	// 	name   string			//в данном случае - не подходит, поскольку нужен именно тот же самый тип.
	// 	age    int32			//[]persons.Person
	// 	weight float64
	// }{
	// 	{name: "Kevin",
	// 		age:    24,
	// 		weight: 68.5,
	// 	},
	// 	{name: "Anna",
	// 		age:    35,
	// 		weight: 70.2,
	// 	},
	// 	{name: "Michael",
	// 		age:    45,
	// 		weight: 90.5,
	// 	},
	// 	{name: "Marina",
	// 		age:    32,
	// 		weight: 60.5,
	// 	},
	// }

	inputTable := []persons.Person{ //используем тип, описанный в пакете persons - для входного слайса.

		{Name: "Kevin",
			Age:    24,
			Weight: 68.5,
		}, {
			Name:   "Anna",
			Age:    35,
			Weight: 70.2,
		}, {
			Name:   "Michael",
			Age:    45,
			Weight: 90.5,
		}, {
			Name:   "Marina",
			Age:    32,
			Weight: 60.5,
		},
	}

	all := persons.GetPersons(inputTable)

	file, err := os.Create("persons.dat")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close() //закрываем файл

	for _, value := range all {

		fmt.Fprintf(file,
			"%-10s %-10d %-10.2f\n", //запись в файл с указанием нужного типа
			//"%s %d %f\n",
			value.Name, value.Age, value.Weight)

	}

}
