package fibonacchi

func Loop(num int) int {

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

func LoopSlice(num int) int {

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

func Recursion(num int) int {

	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	return Recursion(num-1) + Recursion(num-2)
}
