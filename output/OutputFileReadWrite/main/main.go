package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	writeStr()
	writeByte()

	readFromFile()

	fmt.Println("Done!")
}

//writing to file - string
func writeStr() {

	text := "Hello Sam! \n" + "Hello John! \n" + "Hello Amanda! \n"
	file, err := os.Create("helloEveryone.txt")

	if err != nil { //обработка ошибки
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString(text) //для записи текста можно применить метод WriteString

}

//writing to file - byte
func writeByte() {

	data := []byte("Hello Kate!")
	file, err := os.Create("helloKate.bin")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data) //для записи бинарной информации - применяем метод Write (интерфейс - io.Writer)

}

//read from file
func readFromFile() {

	file, err := os.Open("helloEveryone.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data) //читаем порционно
		if err == io.EOF {        //если конец файла - выходим из цикла
			break
		}
		fmt.Print(string(data[:n]))
	}

}
