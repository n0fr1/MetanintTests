package rw

import (
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name   string
	Age    int32
	Weight float64
}

func WriteData(filename string) {

	var people = []Person{
		{"Tom", 24, 68.5},
		{"Bob", 25, 64.2},
		{"Sam", 27, 73.6},
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, p := range people {
		fmt.Fprintf(file, "%s %d %.2f\n", p.Name, p.Age, p.Weight)
	}

}

func ReadData(filename string) []Person {

	var people []Person //дополнительно добавил для возможности теста - чтобы было считано из файла.

	pers := Person{} //сюда считываем

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		_, err = fmt.Fscanf(file, "%s %d %f\n", &pers.Name, &pers.Age, &pers.Weight)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		people = append(people, pers)

		fmt.Printf("%-8s %-8d %-8.2f\n", pers.Name, pers.Age, pers.Weight)

	}

	return people //возвращаем для возможности теста
}
