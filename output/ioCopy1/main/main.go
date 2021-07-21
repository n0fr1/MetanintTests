package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		panic(err)
	}

}
