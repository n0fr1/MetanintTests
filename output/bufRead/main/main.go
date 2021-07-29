package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("../../bufWrite/main/some.dat")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file) //для создания потока вывода через буфер

	for {

		line, err := reader.ReadString('\n') //построчное считывание - пока не обнаружен символ переноса строки

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Print(line)
	}

}
