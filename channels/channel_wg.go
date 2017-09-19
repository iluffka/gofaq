package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	c := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		worker := &Worker{id: i}
		go worker.process(c, wg)
	}

	for {
		c <- rand.Int()
		//time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		data := <-c
		fmt.Printf("обработчик %d получил %d\n", w.id, data)
	}
}