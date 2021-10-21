//пример с работой каналов через обычную функцию.

package main

import (
	"fmt"

	checkinput "github.com/n0fr1/MetanintTests/basics/recursion/checkInput"
)

func main() {

	fmt.Print("Input number for counting factorial: ")
	num := checkinput.GetNum() //проверка ввода пользователя

	intCh := make(chan int)
	go factorial(num, intCh)                          // вызов горутины
	fmt.Printf("%s %d %s", "Result: ", <-intCh, "\n") //получение данных из канала

	fmt.Println("The end")

}

func factorial(n int, ch chan int) {

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result //отправка данных в канал
}

//или данный пример можно переписать, добавив буферизированный канал и убрав горутину.
//Так тоже будет работать:
// intCh := make(chan int, 1)
// factorial (num, intCh)
