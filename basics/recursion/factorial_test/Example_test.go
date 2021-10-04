package factorial_test

import (
	"fmt"

	"github.com/n0fr1/MetanintTests/basics/recursion/factorial"
)

func ExampleLoop() {

	result := factorial.Loop(5)
	fmt.Printf("%d", result)
	//Output: 120

}

func ExampleRecursion() {
	result := factorial.Recursion(8)
	fmt.Printf("%d", result)
	//Output: 40320
}
