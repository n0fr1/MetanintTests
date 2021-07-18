package write_test

import (
	"fmt"

	"github.com/n0fr1/MetanintTests/output/OutputWrite/writing"
)

func ExamplePhoneWriter() {

	bytes1 := []byte("+1(234)567-8900")

	writer := writing.PhoneWriter{}

	lenNum, _ := writer.Write(bytes1)
	bytes1 = bytes1[0:lenNum]

	fmt.Println(string(bytes1))
	//Output: 12345678900

}
