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

	var operation string

	fmt.Print("Введите первое число: ")

	argFirst, err := testInput()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	argSecond, err := testInput()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Print("Введите операцию: ")
	fmt.Scanln(&operation)

	f, err := selectFn(operation)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if argSecond == 0 && operation == "/" {
		fmt.Println("На ноль делить нельзя!")
		return
	}

	fmt.Println(f(argFirst, argSecond))
}

func sum(x, y float64) string {
	return fmt.Sprint(x) + " + " + fmt.Sprint(y) + " = " + fmt.Sprint(x+y)
}

func subtract(x, y float64) string {
	return fmt.Sprint(x) + " - " + fmt.Sprint(y) + " = " + fmt.Sprint(x-y)
}

func multiply(x, y float64) string {
	return fmt.Sprint(x) + " * " + fmt.Sprint(y) + " = " + fmt.Sprint(x*y)
}

func divide(x, y float64) string {
	return fmt.Sprint(x) + " / " + fmt.Sprint(y) + " = " + fmt.Sprint(x/y)
}

func selectFn(operation string) (func(float64, float64) string, error) {

	if operation == "+" {
		return sum, nil
	}

	if operation == "-" {
		return subtract, nil
	}

	if operation == "*" {
		return multiply, nil
	}

	if operation == "/" {
		return divide, nil
	}

	return divide, errors.Wrap(errors.New("Не выбрана операция из доступных (+,-,*,/)\n"), "Ошибка") //оборачиваем ошибку
}

func testInput() (float64, error) {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n') //ReadString - если все, что нужно ввод одного слова. Single word
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //strings.TrimSpace - убирает пробел
	if err != nil {
		return 0, errors.Wrap(errors.New("Вы ввели не число! \n"), "Ошибка")
	}

	return num, err
}
