package persons

type AllPersons []Person

type Person struct {
	Name   string //чтобы структура была доступна в других модулях, названия полей - должны быть
	Age    int32  //с большой буквы
	Weight float64
}

func GetPersons(inputTable AllPersons) AllPersons {

	//var people AllPersons
	people := AllPersons{} //есть разница между двумя вариантами создания. В первом случае - переменная равна nil и тест не проходим.

	for _, value := range inputTable {

		if value.Age <= 40 {
			strPerson := fillstruct(value)      //надуманный пример.
			people = append(people, *strPerson) //допустим, нужно получить слайс с определенными сотрудниками и его потом вывести в файл
		}
	}

	return people
}

func fillstruct(someone Person) *Person {

	return &Person{
		Name:   someone.Name,
		Age:    someone.Age,
		Weight: someone.Weight,
	}
}

//хотелось бы добавить в этот пример интерфейс.
