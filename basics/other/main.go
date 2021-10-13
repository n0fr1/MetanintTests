//Задача: есть возрастающий числовой массив.
//Необходимо его разбить на диапазоны, на которых он непрерывен
//Например: 1,3,4,5,7,8,9,10
//должен быть разбит: 1-1, 3-5,7-10

package main

import "fmt"

func main() {
	var1()
}

func var1() {

	var result []int

	slice := []int{1, 3, 4, 6, 7, 8, 9, 10, 11, 15, 16, 20, 24, 25, 26, 40, 41, 42, 44, 50}

	for ind, val := range slice {

		if ind == 0 { //первый элемент
			result = append(result, val)
			continue
		}

		if val-slice[ind-1] > 1 { //разрыв
			result = append(result, slice[ind-1])
			result = append(result, val)
		}

		if ind == len(slice)-1 { //последний элемент
			result = append(result, val)
		}
	}

	fmt.Println(result)
}
