package factorial_test

import (
	"reflect"
	"testing"

	"github.com/n0fr1/MetanintTests/basics/recursion/factorial"
)

func TestRecursion(t *testing.T) {

	in := 5
	out := 120

	t.Run("Testing", func(t *testing.T) {
		result := factorial.Recursion(in)

		if !reflect.DeepEqual(out, result) {
			t.Error("Expected", out, " got ", result)
		}

	})
}

func TestLoop(t *testing.T) {

	in := 5
	out := 120

	t.Run("Testing", func(t *testing.T) {
		result := factorial.Loop(in)

		if !reflect.DeepEqual(out, result) {
			t.Error("Expected", out, " got ", result)
		}

	})
}
