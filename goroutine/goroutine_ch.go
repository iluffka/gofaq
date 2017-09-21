package main

import (
	"sync"
	"fmt"
)

var wg = sync.WaitGroup{}


/*
func main()  {
	ch := make(chan int, 5)
	wg.Add(2)
	go func(ch <-chan int) {   //канал ch только ПРИНИМАЕТ сообщения
		//defer close(ch) deadlock!
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { //канал ch только отправляет сообщения
		//defer close(ch)
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}*/

/*
func main()  {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {   //канал ch только ПРИНИМАЕТ сообщения
		//defer close(ch) deadlock!
		for i := range ch {
			fmt.Println("i: ", i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) { //канал ch только отправляет сообщения
		defer close(ch)
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}*/

func main()  {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <- ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		defer close(ch)
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}