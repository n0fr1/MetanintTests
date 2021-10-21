//сколько итераций пройдет по циклу?
//3 итерации

package main

import "fmt"

func main() {

	b := 3
	for a := 1; b > a; a++ {
		if a == 2 { //дополнительная итерация
			b += 1
		}
		fmt.Println(a)
	}

}
