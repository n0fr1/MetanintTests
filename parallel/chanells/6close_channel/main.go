package main

import "fmt"

func main() {

	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 3
	close(intCh)

	//intCh <- 24 //ошибка - канал уже закрыт.

	fmt.Println(<-intCh) //10
	fmt.Println(<-intCh) //3
	fmt.Println(<-intCh) //0

}

//После закрытия канала - мы не можем послать в него новые данные.
//Однако можно получить ранее добавленные данные.
//При попытке получить данные, которых - нет. Получим значение - ноль.
