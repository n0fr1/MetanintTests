package main

import (
	"fmt"
	"os"
)

func read() (int, []byte) {

	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	buffer := make([]byte, 64)

	var n int
	n, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return n, buffer
}

func main() {
	end, data := read()
	data = data[0:end]
	fmt.Println(string(data))
}
