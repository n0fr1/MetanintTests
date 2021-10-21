//проверка состояния канала.

package main

import "fmt"

func main() {
	intCh := make(chan int, 3)

	intCh <- 10
	intCh <- 3
	close(intCh) //канал закрыт

	for i := 0; i < cap(intCh); i++ {
		if val, opened := <-intCh; opened { //если opened - true, канал открыт.
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed")
		}
	}
}

//другой вариант записи, более привычный:
//val,opened := <-intCh
//if opened {
//	fmt.Println(val)
//} else {
//	fmt.Pritln("Channel closed")
//}
