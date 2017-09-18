package main

import (
	"fmt"
	"sync"
)

/*func main()  {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(i)
		}(i, wg)
	}
	wg.Wait()
}*/


//Next example

var (
	wg = sync.WaitGroup{}
	counter = 0
)

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello()  {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment()  {
	counter++
	wg.Done()
}