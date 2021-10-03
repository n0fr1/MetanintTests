package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	calcFactorial() //расчет факториала - через цикл и рекурсию
	calcFibonachi() //вывод чисел фибоначчи
}

func calcFactorial() {

	fmt.Print("Введите факториал: ")
	num := getNum()

	result := getFactorialRec(num)
	fmt.Printf("%s %d %s", "Факториал рекурсия:", result, "\n")

	result = getFactorialLoop(num)
	fmt.Printf("%s %d %s", "Фактория цикл:", result, "\n")

}

func getFactorialRec(x int) int {

	if x == 0 {
		return 1
	}

	return x * getFactorialRec(x-1)

}

func getFactorialLoop(x int) int {

	if x == 0 {
		return 1
	}

	result := 1

	for count := x; count > 0; count-- {
		result *= count
	}

	return result

}

func calcFibonachi() {

	fmt.Print("Введите номер числа фибоначчи: ")
	num := getNum()

	result := getFiboRec(num)
	fmt.Printf("%s %d %s", "Фибоначчи рекурсия:", result, "\n")

	result = getFiboLoop(num)
	fmt.Printf("%s %d %s", "Фибоначчи цикл:", result, "\n")

	result = getFiboLoopSlice(num)
	fmt.Printf("%s %d %s", "Фибоначчи цикл (через слайс):", result, "\n")

}

func getFiboLoop(num int) int {

	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	result := 0
	a := 0
	b := 0

	for x := 0; x < num-1; x++ {

		if a == 0 && b == 0 { //первый проход
			a = x
			b = x + 1
		}

		result = a + b
		b = a
		a = result

	}

	return result
}

func getFiboLoopSlice(num int) int {

	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	var slice = make([]int, num)

	slice[0] = 0
	slice[1] = 1

	for x := 2; x < num; x++ {
		slice[x] = slice[x-1] + slice[x-2]
	}

	return slice[num-1]
}

func getFiboRec(num int) int {

	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	return getFiboRec(num-1) + getFiboRec(num-2)
}

func getNum() int {

	num, err := testInput()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	return num
}

func testInput() (int, error) {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		return num, errors.Wrap(errors.New("Введен символ! \n"), "Ошибка")
	}

	if num < 0 {
		return num, errors.Wrap(errors.New("Введено отрицательно число! \n"), "Ошибка")
	}

	return num, err
}
