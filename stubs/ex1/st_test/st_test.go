package st_test

import (
	"testing"

	"github.com/n0fr1/MetanintTests/stubs/ex1/gettingPhone"
)

//стаб. Тестим функцию find, подменяя search
type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*gettingPhone.Someone, firstName, lastName string) *gettingPhone.Someone {
	return &gettingPhone.Someone{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestStub(t *testing.T) {

	annaD := &gettingPhone.Someone{
		FirstName: "Anna",
		LastName:  "Dobreva",
		Phone:     "11111"}

	janeQ := &gettingPhone.Someone{
		FirstName: "Jane",
		LastName:  "Qwerty",
		Phone:     "22222"}

	simonK := &gettingPhone.Someone{
		FirstName: "Simon",
		LastName:  "Kollin",
		Phone:     "33333"}

	people := []*gettingPhone.Someone{annaD, janeQ, simonK}

	phonebook := gettingPhone.Phonebook{
		People: people,
	}

	fakePhone := "+31 65 222 333"
	phone, _ := phonebook.Find(StubSearcher{fakePhone}, "Anna", "Dobreva")

	if phone != fakePhone {
		t.Errorf("Want '%s', go '%s'", fakePhone, phone)
	}

}
