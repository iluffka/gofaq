package main

import (
	"fmt"
	"sync"
)

func main() {
	inChan := make(chan []string, )
	outChan := make(chan []string)
	buf := make(chan struct{}, 3) //канал для отслеживания гоурутин, в данном случае 3

	go getData(inChan)
	go processData(inChan, outChan, buf)
	printData(outChan)
}

func getData(out chan []string) {
	defer close(out)
	//получаем входные данные
	out <- []string{"a1", "b"}
	out <- []string{"a2", "b"}
	out <- []string{"a3", "b"}
	out <- []string{"a4", "b"}
	out <- []string{"a5", "b"}
}

func processData(in chan []string, out chan []string, buf chan struct{}) {
	defer close(out)

	wg := new(sync.WaitGroup)

	for d := range in {
		wg.Add(1)
		buf <- struct{}{}
		go func(data []string) {
			defer wg.Done()
			out <- append(data, "c")
			<-buf
		}(d)
	}
	wg.Wait()
}


func printData(in chan []string) {
	for data := range in {
		fmt.Println(data)
	}
}