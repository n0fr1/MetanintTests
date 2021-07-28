package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	file, err := os.Create("confeve.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	tTime := time.Date(2021, time.June, 30, 8, 5, 2, 0, time.Local)

	_, err = fmt.Fprintf(file, "%s %s\n", "Сегодня ", tTime.Format("02-01-2006")) //число используемое для форматирования
	if err != nil {                                                               //обработка ошибки
		panic(err)
	}

	_, err = fmt.Fprintf(file, "%s", "Хорошая погода!")
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintln(os.Stdout, "Good weather!")
	if err != nil {
		panic(err)
	}
}
