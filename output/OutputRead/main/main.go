package main

import (
	"fmt"

	"github.com/n0fr1/MetanintTests/output/OutputRead/reading"
)

func main() {

	phone1 := reading.PhoneReader("+1(234)567 9010")
	phone2 := reading.PhoneReader("+2(345)777 8090")

	buffer := make([]byte, len(phone1)) //срез размером с номер телефона
	lenNum, _ := phone1.Read(buffer)    //передаем в метод Read
	buffer = buffer[0:lenNum]           //отсекаем лишние байты
	fmt.Printf("%s \n", string(buffer))

	buffer = make([]byte, len(phone2))
	lenNum, _ = phone2.Read(buffer)
	buffer = buffer[0:lenNum]
	fmt.Printf("%s \n", string(buffer))
}
