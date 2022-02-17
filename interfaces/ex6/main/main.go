//пример с несколькими полями в структуре
package main

import "fmt"

type Show interface {
	infoAuto()
}

type Auto struct {
	name  string
	model string
}

func (s *Auto) infoAuto() {
	fmt.Printf("%s %s %s %s \n", "name:", s.name, "model:", s.model)
}

func main() {

	getModel := &Auto{
		name:  "Lada",
		model: "largus",
	}

	getModel.infoAuto()
}
