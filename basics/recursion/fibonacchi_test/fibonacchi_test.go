package fibonacchi_test

import (
	"reflect"
	"testing"

	"github.com/n0fr1/MetanintTests/basics/recursion/fibonacchi"
	"github.com/stretchr/testify/assert"
)

func TestLoopAssert(t *testing.T) {

	testCases := []struct {
		name string
		n    int
		want int
	}{
		{name: "zero",
			n:    0,
			want: 0,
		},
		{name: "one",
			n:    1,
			want: 1,
		},
		{name: "seven",
			n:    7,
			want: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fibonacchi.Loop(tc.n))
		})
	}
}

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
