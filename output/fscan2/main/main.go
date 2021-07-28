package main

import "github.com/n0fr1/MetanintTests/output/fscan2/rw"

func main() {

	filename := "people.dat"
	rw.WriteData(filename)
	rw.ReadData(filename)
}
