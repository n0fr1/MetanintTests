package readtest_test

import (
	"fmt"

	"github.com/n0fr1/MetanintTests/output/OutputRead/reading"
)

func ExamplePhoneReader() {

	phone1 := reading.PhoneReader("+1(234)567 9010")

	buffer := make([]byte, len(phone1)) //побайтовый срез достаточной длины
	lenNum, _ := phone1.Read(buffer)    //передаем в метод Read
	buffer = buffer[0:lenNum]

	fmt.Printf("%s", string(buffer))
	//Output: 12345679010
}
