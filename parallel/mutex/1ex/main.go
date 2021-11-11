package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func main() {

	var mutex sync.Mutex
	ch := make(chan bool)

	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("The end")
}

func work(number int, ch chan bool, mutex *sync.Mutex) {

	mutex.Lock()
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock()

	ch <- true
}
