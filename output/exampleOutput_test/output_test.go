package exampleOutput_test

import (
	"fmt"

	"github.com/n0fr1/MetanintTests/persons"
)

//ExampleT is example test
func ExamplePerson() { //Example(Person) - в скобках обязательное название структуры

	in := []persons.Person{ //используем тип, описанный в пакете persons - для входного слайса.

		{Name: "Kevin",
			Age:    24,
			Weight: 68.5,
		}, {
			Name:   "Michael",
			Age:    45,
			Weight: 90.5,
		},
	}

	fmt.Println(persons.GetPersons(in))
	//Output: [{Kevin 24 68.5}]
}
