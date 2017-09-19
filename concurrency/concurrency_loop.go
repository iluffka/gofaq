package main

import (
	"sync"
	"fmt"
)

func main() {
	inChan := make(chan []string, 100)
	outChan := make(chan []string, 100)

	go getIn(inChan)
	go fillOut(inChan, outChan)
	saveOut(outChan)
}


func getIn(out chan []string) {
	defer close(out)

	// get some input data to in
	out <- []string{"a1", "b"}
	out <- []string{"a2", "b"}
	out <- []string{"a3", "b"}
	out <- []string{"a4", "b"}
	out <- []string{"a5", "b"}
}

func fillOut(in chan []string, out chan []string) {
	defer close(out)

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go subFillOut(in, out, &wg)
	}
	wg.Wait()
}

func subFillOut(in chan []string, out chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range in {
		out <- append(data, "c")
	}
}

func saveOut(in chan []string) {
	for data := range in {
		fmt.Println(data)
	}
}
