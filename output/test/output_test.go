package test

import "testing"

func TestOutput(t *testing.T) {

	cases := []struct {
		name   string
		age    int32
		weight float64
	}{
		{name: "Kevin",
			age:    24,
			weight: 68.5,
		},
		{name: "Anna",
			age:    35,
			weight: 70.2,
		},
		{name: "Marina",
			age:    32,
			weight: 60.5,
		},
	}

	for _, tc = range cases {
		t.Run(tc.name, func(t *testing.T){
			//result := 
		}
	}
}
