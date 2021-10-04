package main

import (
	"fmt"

	checkinput "github.com/n0fr1/MetanintTests/basics/recursion/checkInput"
	"github.com/n0fr1/MetanintTests/basics/recursion/factorial"
	"github.com/n0fr1/MetanintTests/basics/recursion/fibonacchi"
)

func main() {
	calcFactorial() //расчет факториала - через цикл и рекурсию
	calcFibonachi() //вывод чисел фибоначчи
}

func calcFactorial() {

	fmt.Print("Введите факториал: ")
	num := checkinput.GetNum()
	result := factorial.Recursion(num)
	fmt.Printf("%s %d %s", "Факториал рекурсия:", result, "\n")

	result = factorial.Loop(num)
	fmt.Printf("%s %d %s", "Фактория цикл:", result, "\n")

}

func calcFibonachi() {

	fmt.Print("Введите номер числа фибоначчи: ")
	num := checkinput.GetNum()

	result := fibonacchi.Recursion(num)
	fmt.Printf("%s %d %s", "Фибоначчи рекурсия:", result, "\n")

	result = fibonacchi.Loop(num)
	fmt.Printf("%s %d %s", "Фибоначчи цикл:", result, "\n")

	result = fibonacchi.LoopSlice(num)
	fmt.Printf("%s %d %s", "Фибоначчи цикл (через слайс):", result, "\n")

}
