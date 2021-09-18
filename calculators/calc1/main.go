package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var operation string

	fmt.Print("Введите первое число: ")
	argFirst, err := testNumber() //проверяем 1-й введенный аргумент на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	argSecond, err := testNumber() //проверяем 2-й введенный аргумент на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Print("Введите операцию: ")
	fmt.Scanln(&operation)

	if !(argSecond == 0 && operation == "/") {
		view(&argFirst, &argSecond, &operation)
	} else {
		fmt.Print("На ноль делить нельзя!")
	}

}

func view(argFirst *float64, argSecond *float64, calcOp *string) {

	result, err := count(argFirst, argSecond, calcOp) //проверка на ошибку ввода знака операции
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	//здесь немного "поигрался" с выводом.
	if *argFirst == math.Round(*argFirst) && *argSecond == math.Round(*argSecond) { //если оба числа "равны" округленным выражениям
		fmt.Printf("%.0f %s %.0f %s %.0f\n", *argFirst, *calcOp, *argSecond, "=", result) //то тогда нет смысла показывать нули после "запятой"
	} else if *argFirst == math.Round(+*argFirst) && *argSecond != math.Round(*argSecond) {
		fmt.Printf("%.0f %s %f %s %f\n", *argFirst, *calcOp, *argSecond, "=", result)
	} else {
		fmt.Printf("%f %s %f %s %f\n", *argFirst, *calcOp, *argSecond, "=", result)
	}

}

func count(argFirst *float64, argSecond *float64, calcOp *string) (float64, error) {

	switch *calcOp {
	case "+":
		return *argFirst + *argSecond, nil
	case "-":
		return *argFirst - *argSecond, nil
	case "*":
		return *argFirst * *argSecond, nil
	case "/":
		return *argFirst / *argSecond, nil
	default:
		return 0, errors.New("Не выбрана операция из доступных: * / + - ")
	}

}

func testNumber() (float64, error) {

	input := bufio.NewScanner(os.Stdin) //NewScanner - можно использоваться в цикле, если нужно считать "много слов". Multiply words
	input.Scan()

	num, err := strconv.ParseFloat(input.Text(), 64)
	return num, err
}
