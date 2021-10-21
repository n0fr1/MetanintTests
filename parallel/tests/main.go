package main

import "fmt"

func main() {
	intCh := make(chan int, 1)
	intCh <- 10
	fmt.Println(<-intCh)
}
