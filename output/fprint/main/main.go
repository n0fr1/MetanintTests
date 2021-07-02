package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	file, err := os.Create("confeve.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tTime := time.Date(2021, time.June, 30, 8, 5, 2, 0, time.Local)

	fmt.Fprintf(file, "%s %s\n", "Сегодня ", tTime.Format("02-01-2006")) //число используемое для форматирования
	fmt.Fprintf(file, "%s", "хорошая погода")

}
