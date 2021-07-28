package r_test

import (
	"reflect"
	"testing"

	"github.com/n0fr1/MetanintTests/output/fscan2/rw"
)

//тест к форматированному выводу.
func TestPerson(t *testing.T) {

	var people = []rw.Person{
		{"Tom", 24, 68.5},
		{"Bob", 25, 64.2},
		{"Sam", 27, 73.6},
	}

	t.Run("Testing", func(t *testing.T) {
		result := rw.ReadData("../main/people.dat")

		for index := range people {
			if !reflect.DeepEqual(people[index], result[index]) {
				t.Error("Expected", people[index], " got ", result[index])
			}
		}
	})
}
