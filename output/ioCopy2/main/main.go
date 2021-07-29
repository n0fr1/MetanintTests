package main

import (
	"fmt"
	"io"
	"os"

	"github.com/n0fr1/MetanintTests/output/ioCopy2/phones"
)

func main() {
	phone1 := phones.PhoneReader("+1(234)567 90-10")

	_, err := io.Copy(os.Stdout, phone1)
	if err != nil {
		panic(err)
	}
	fmt.Println()

}
