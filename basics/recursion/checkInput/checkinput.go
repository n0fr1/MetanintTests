//тест на ввод только целых неотрицательных чисел

package checkinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func GetNum() int {

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
