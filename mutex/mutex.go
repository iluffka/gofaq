package main

import (
	"sync"
	"fmt"
)

var (
	wg = sync.WaitGroup{}
	counter = 0
	m = sync.Mutex{}
)

func main()  {
	//fmt.Printf("Thread %v\n", runtime.GOMAXPROCS(-1))
	var i int
	for i = 0; i < 10; i++ {
		wg.Add(2)
		m.Lock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello()  {
	fmt.Printf("Hello #%v\n", counter)
	m.Unlock()
	wg.Done()
}

func increment()  {
	counter++
	m.Unlock()
	wg.Done()
}