package factorial

func Recursion(x int) int {

	if x == 0 {
		return 1
	}

	return x * Recursion(x-1)

}

func Loop(x int) int {

	if x == 0 {
		return 1
	}

	result := 1

	for count := x; count > 0; count-- {
		result *= count
	}

	return result

}
