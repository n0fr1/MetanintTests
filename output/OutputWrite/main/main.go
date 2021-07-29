package main

import (
	"fmt"

	"github.com/n0fr1/MetanintTests/output/OutputWrite/writing"
)

func main() {

	bytes1 := []byte("+1(234)567-8900")
	bytes2 := []byte("+2-352-889-777")

	writer := writing.PhoneWriter{}

	lenNum, _ := writer.Write(bytes1)
	bytes1 = bytes1[0:lenNum]

	fmt.Println(string(bytes1))

	lenNum, _ = writer.Write(bytes2)
	bytes2 = bytes2[0:lenNum]

	fmt.Println(string(bytes2))

}
