//Пример из маленькой книжки GO
//Буферизированные каналы.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func main() {

	c := make(chan int)

	for i := 0; i < 5; i++ { //сначала запускается несколько обработчиков
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for { //затем даем им немного поработать
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}

}

func (w *Worker) process(c chan int) {

	for {
		data := <-c
		fmt.Printf("обработчик %d получил %d\n", w.id, data)
	}
}
