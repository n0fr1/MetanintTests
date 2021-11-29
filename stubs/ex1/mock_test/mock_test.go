package mock_test

import (
	"testing"

	"github.com/n0fr1/MetanintTests/stubs/ex1/gettingPhone"
)

//мок тест. Проверяем также, что и вызывается метод Search
type MockSearcher struct {
	phone         string
	methodsToCall map[string]bool
}

func (ms *MockSearcher) Search(people []*gettingPhone.Someone, firstName, lastName string) *gettingPhone.Someone {
	ms.methodsToCall["Search"] = true
	return &gettingPhone.Someone{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}

	ms.methodsToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {

	for methodName, called := range ms.methodsToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't.", methodName)
		}
	}

}

func TestFindCallsSearchAndReturnsPersonUsingMock(t *testing.T) {

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

	fakePhone := "+31 65 222 333"

	phonebook := gettingPhone.Phonebook{
		People: people,
	}

	mock := &MockSearcher{phone: fakePhone}
	mock.ExpectToCall("Search")

	phone, _ := phonebook.Find(mock, "Anna", "Dobreva")

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}

	mock.Verify(t)
}
