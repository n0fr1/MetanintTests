package fibonacchi_test

import (
	"reflect"
	"testing"

	"github.com/n0fr1/MetanintTests/basics/recursion/fibonacchi"
)

func TestLoop(t *testing.T) {

	in := 9
	out := 21

	t.Run("Testing", func(t *testing.T) {
		result := fibonacchi.Loop(in)

		if !reflect.DeepEqual(out, result) {
			t.Error("Expected", out, " got ", result)
		}

	})
}

func TestLoopSlice(t *testing.T) {

	in := 7
	out := 8

	t.Run("Testing", func(t *testing.T) {
		result := fibonacchi.LoopSlice(in)

		if !reflect.DeepEqual(out, result) {
			t.Error("Expected", out, " got ", result)
		}

	})

}

func TestRecursion(t *testing.T) {

	in := 10
	out := 55

	t.Run("Testing", func(t *testing.T) {
		result := fibonacchi.Recursion(in)

		if !reflect.DeepEqual(out, result) {
			t.Error("Expected", out, " got ", result)
		}

	})

}
