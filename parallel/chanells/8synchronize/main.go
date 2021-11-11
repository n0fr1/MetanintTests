package main

import (
	"fmt"

	checkinput "github.com/n0fr1/MetanintTests/basics/recursion/checkInput"
)

func main() {

	structCh := make(chan struct{}) //используем пустую структуру только для того, чтобы синхронизировать объект
	num := checkinput.GetNum()
	fmt.Printf("%s", "\n")

	results := make([]int, num+1)
	go factorial(num, structCh, results)

	<-structCh //ожидаем закрытия канала structCh

	for i, v := range results {
		fmt.Println(i, " - ", v)
	}
}

func factorial(n int, ch chan struct{}, results []int) {

	defer close(ch) //без записи - закрываем канал

	result := 1

	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}

}
