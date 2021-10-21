//Если необходимо определить функцию, которая возвращает канал
//то все операции чтения-записи в канал - следует вынести в отдельную горутину
//иначе получаем - блокировку.

package main

import "fmt"

func main() {
	fmt.Println("Start")
	//создание канала и получение

	fmt.Println(<-createChan(5))
	fmt.Println("End")

}

func createChan(n int) chan int {
	ch := make(chan int) //создаем канал

	go func() {
		ch <- n
	}()

	return ch
}
