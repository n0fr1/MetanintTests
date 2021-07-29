package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	fmt.Fprintf(file, "%s", "Hello!")
	fmt.Println(file.Name())

}
