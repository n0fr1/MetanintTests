package fibonacchi_test

import (
	"fmt"

	"github.com/n0fr1/MetanintTests/basics/recursion/fibonacchi"
)

func ExampleLoop() {

	result := fibonacchi.Loop(7)
	fmt.Printf("%d", result)
	//Output: 8

}

func ExampleLoopSlice() {
	result := fibonacchi.LoopSlice(9)
	fmt.Printf("%d", result)
	//Output: 21
}

func ExampleRecursion() {
	result := fibonacchi.Recursion(10)
	fmt.Printf("%d", result)
	//Output: 55
}
