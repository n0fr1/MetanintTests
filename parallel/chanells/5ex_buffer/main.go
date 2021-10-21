//Можно определить канал, как доступный только для отправки данных
//или только для получения данных

package main

import (
	"fmt"

	checkinput "github.com/n0fr1/MetanintTests/basics/recursion/checkInput"
)

func main() {

	fmt.Print("Input number for counting factorial: ")
	num := checkinput.GetNum() //проверка ввода пользователя

	intCh := make(chan int, 1)

	go factorial(num, intCh)

	fmt.Printf("%s %d %s", "Result: ", <-intCh, "\n") //получение данных из канала
	fmt.Println("The end")

}

func factorial(n int, ch chan<- int) { //канал, доступный только для отправки данных.

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result //Внутри функции factorial мы можем только отправлять данные в канал, получать не можем.

}
