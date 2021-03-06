package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	rows := []string{
		"Hello Go!",
		"Welcome to Golang",
	}

	file, err := os.Create("some.dat")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	writer := bufio.NewWriter(file) //создание потока вывода через буфер

	for _, row := range rows {
		writer.WriteString(row)  // запись строки
		writer.WriteString("\n") // перевод строки
	}

	writer.Flush() //сбрасываем данные из буфера в файл
}
