// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var counter = 0

// func main() {

// 	var mutex sync.Mutex

// 	for i := 0; i < 2; i++ {
// 		go incr(&mutex)
// 	}

// 	time.Sleep(time.Millisecond * 10)
// }

// func incr(mutex *sync.Mutex) {
// 	mutex.Lock()
// 	counter++
// 	mutex.Unlock()

// 	fmt.Println(counter)
// }

//deadlock
package main

import (
	"sync"
	"time"
)

var (
	lock sync.Mutex
)

func main() {
	go func() { lock.Lock() }()
	time.Sleep(time.Millisecond * 10)
	lock.Lock()

}
