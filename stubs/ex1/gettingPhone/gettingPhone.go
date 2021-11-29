//доработать пример, чтобы здесь был слайс с фамилиями

package gettingPhone

import (
	"errors"
)

var (
	ErrMissingArgs   = errors.New("firstName and LastName are mandatory arguments")
	ErrNoPersonFound = errors.New("no person found")
)

type Someone struct {
	FirstName string
	LastName  string
	Phone     string
}

type Searcher interface {
	Search(people []*Someone, firstName string, lastname string) *Someone
}

type Phonebook struct {
	People []*Someone
}

func (p *Phonebook) Find(searcher Searcher, firstName, lastname string) (string, error) {

	if firstName == "" || lastname == "" {
		return "", ErrMissingArgs
	}

	person := searcher.Search(p.People, firstName, lastname)

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil

}

func (p *Phonebook) Search(people []*Someone, firstName, lastname string) *Someone {

	for _, val := range people {
		if val.FirstName == firstName && val.LastName == lastname {
			return val
		}
	}

	return nil
}

// func main() {

// 	annaD := &Someone{
// 		FirstName: "Anna",
// 		LastName:  "Dobreva",
// 		Phone:     "11111"}

// 	janeQ := &Someone{
// 		FirstName: "Jane",
// 		LastName:  "Qwerty",
// 		Phone:     "22222"}

// 	simonK := &Someone{
// 		FirstName: "Simon",
// 		LastName:  "Kollin",
// 		Phone:     "33333"}

// 	people := []*Someone{annaD, janeQ, simonK}

// 	phonebook := Phonebook{
// 		People: people,
// 	}

// 	result, err := phonebook.Find(&phonebook, "Jack", "Orin")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}

// }
