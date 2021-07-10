package exampleOutput_test

import (
	"reflect"
	"testing"

	"github.com/n0fr1/MetanintTests/persons"
)

func TestPerson(t *testing.T) {

	in := persons.AllPersons{ //используем тип, описанный в пакете persons - для входного слайса.

		{
			Name:   "Michael",
			Age:    45,
			Weight: 90.5,
		},
	}

	out := persons.AllPersons{} //используем тип, описанный в пакете persons - для входного слайса.

	t.Run("Testing", func(t *testing.T) {
		result := persons.GetPersons(in)

		if !reflect.DeepEqual(out, result) {
			t.Error("Expected", out, " got ", result)
		}
	})
}
