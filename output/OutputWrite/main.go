package main

import "fmt"

type phoneWriter struct{}

func main() {

	bytes1 := []byte("+1(234)567-8900")
	bytes2 := []byte("+2-352-889-777")

	writer := phoneWriter{}
	writer.Write(bytes1)
	writer.Write(bytes2)

}

func (p phoneWriter) Write(bs []byte) (int, error) {

	if len(bs) == 0 {
		return 0, nil
	}

	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]))
		}
	}
	fmt.Println()

	return len(bs), nil

}
