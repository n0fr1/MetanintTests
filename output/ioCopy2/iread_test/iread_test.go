package iread_test

import (
	"bytes"
	"fmt"
	"io"

	"github.com/n0fr1/MetanintTests/output/ioCopy2/phones"
)

func ExamplePhoneReader() {

	phone1 := phones.PhoneReader("+1(234)567 90-10")
	buffer := &bytes.Buffer{}

	_, err := io.Copy(buffer, phone1)
	if err != nil {
		panic(err)
	}

	fmt.Println(buffer.String())
	//Output: 12345679010
}
