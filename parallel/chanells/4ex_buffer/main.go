//Буферизированные каналы создаются с помощью функции make().

package main

import "fmt"

func main() {

	intCh := make(chan int, 4) //канал из трех элементов.
	intCh <- 10
	intCh <- 3
	intCh <- 15

	fmt.Printf("%s %d %s %d %s", "capacity:", cap(intCh), "length: ", len(intCh), "\n") //получаем емкость и количество элементов в канале.

	fmt.Printf("%s", "results:\n")
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)

	//если отправить больше элементов, чем вместимость канала - получим deadlock

}
