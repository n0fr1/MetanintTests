package main

import (
	"fmt"
)

func main() {
	indefeniedParams()
	simpleCalc()
	sumNumbers()
	simpleCalcAnonym()
	simpleCalcAnonym2()
	simpleAnonym()
	simpleAnonym2()
	simpleAnonym3()
	simpleAnonym4()
	simpleAnonym5()

}

//пример №1 - с передачей неопределенного количества параметров (одного типа)
func indefeniedParams() {
	add(1, 2, 3)           //sum = 6
	add([]int{1, 2, 3}...) //sum = 6
	add(1, 2, 3, 4)        //sum = 10
}

func add(numbers ...int) {
	var sum = 0
	for _, value := range numbers {
		sum += value
	}

	fmt.Println("sum = ", sum)
}

//пример №2 - передача функции в качестве параметра в другую функцию
func simpleCalc() {
	action(3, 4, plus)
	action(3, 4, multiply)
}

func action(x int, y int, operation func(int, int) string) {
	result := operation(x, y) //plus,multiply - передаем в operation
	fmt.Println(result)
}

//пример №3 - передача функции в качестве параметра в другую функцию
func sumNumbers() {

	slice := []int{-2, 4, 3, -1, 7, -4, 23}
	sumOfEvens := sum(slice, isEven) //сумма четных чисел
	fmt.Println(sumOfEvens)          // -2

	sumOfPositives := sum(slice, isPositive) //сумма положительных чисел
	fmt.Println(sumOfPositives)              //37

}

func sum(numbers []int, criteria func(int) bool) int {

	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}

	return result
}

func isEven(n int) bool {
	return n%2 == 0 //с делением без остатка
}

func isPositive(n int) bool {
	return n > 0
}

//пример №4 - Тоже самое, но через анонимную функцию.
func simpleCalcAnonym() {
	f := actionAn("+")
	fmt.Println(f(3, 4))

	f = actionAn("*")
	fmt.Println(f(3, 4))
}

func actionAn(action string) func(int, int) string {

	if action == "+" { //здесь получаем только саму функцию
		return plus
	}

	return multiply
}

//пример №5 - Другая запись через анонимную функцию.
func simpleCalcAnonym2() {
	var f func(int, int) string = plus
	fmt.Println(f(3, 4))

	f = multiply
	var x = f(3, 4)
	fmt.Println(x)
}

//Общие функции для примеров
func plus(x, y int) string {
	return fmt.Sprint(x) + " + " + fmt.Sprint(y) + " = " + fmt.Sprint(x+y)
}

func multiply(x, y int) string {
	return fmt.Sprint(x) + " * " + fmt.Sprint(y) + " = " + fmt.Sprint(x*y)
}

//пример №6 - Простой пример анонимной функции
func simpleAnonym() {
	f := func(x, y int) int { return x + y }
	fmt.Println(f(3, 4))
	fmt.Println(f(6, 8))
}

//пример №7 - Другой простой пример анонимной функции
func simpleAnonym2() {
	action2(10, 25, func(x, y int) int { return x + y })
	action2(5, 6, func(x, y int) int { return x + y })
}

func action2(n1, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

//пример №8 - Анонимная функция как результат функции
func simpleAnonym3() {
	f := selectFn(1)
	fmt.Println(f(2, 3))

	f = selectFn(1)
	fmt.Println(f(5, 2))

}

func selectFn(n int) func(x, y int) int {

	if n == 1 {
		return func(x, y int) int { return x + y }
	} else if n == 2 {
		return func(x, y int) int { return x - y }
	} else {
		return func(x, y int) int { return x * y }
	}

}

//пример №9 - доступ к окружению, в котором вызываются анонимные функции
func simpleAnonym4() {
	f := square()
	fmt.Println(f()) //9
	fmt.Println(f()) //16
	fmt.Println(f()) //25 - результат разный. кэшируется?
}

//пример №10 - доступ к окружению, в котором вызываются анонимные функции
func simpleAnonym5() {
	f := square()
	fmt.Println(f()) //9

	f = square()
	fmt.Println(f()) //9
}

func square() func() int {
	x := 2

	return func() int {
		x++
		return x * x
	}

}
