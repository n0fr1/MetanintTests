//пустые интерфейсы. Могут использоваться, когда - заранее неизвестно, что именно передаем.
package main

import (
	"fmt"
	"reflect"
)

func main() {

	m := map[string]interface{}{} //ключ - строка, а значение - всегда варьируется, т.е заранее неизвестно.
	m["one"] = 1
	m["two"] = 2.0
	m["three"] = true

	for k, v := range m {
		switch v.(type) {
		case int:
			fmt.Printf("%s is an integer\n", k)
		case float64:
			fmt.Printf("%s is a float64\n", k)
		default:
			fmt.Printf("%s is %v\n", k, reflect.TypeOf(v))

		}
	}
}
