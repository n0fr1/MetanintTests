package main

import (
	"fmt"

	"github.com/n0fr1/MetanintTests/output/OutputRead/reading"
)

func main() {

	phone1 := reading.PhoneReader("+1(234)567 9010")
	phone2 := reading.PhoneReader("+2(345)777 8090")

	buffer := make([]byte, len(phone1)) //побайтовый срез достаточной длины
	phone1.Read(buffer)                 //передаем в метод Read
	fmt.Println(string(buffer))

	buffer = make([]byte, len(phone2))
	phone2.Read(buffer)
	fmt.Println(string(buffer))

}
