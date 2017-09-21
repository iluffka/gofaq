package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

func main() {
	inChan := make(chan []string)
	outChan := make(chan []string)
	countChan := make(chan int)
	buf := make(chan struct{}, 3) //канал для отслеживания гоурутин, в данном случае 3

	go getData(inChan)
	go processData(inChan, outChan, buf, countChan)
	printData(outChan, countChan)
}

func getData(out chan []string) {
	defer close(out)
	//получаем входные данные
	/*out <- []string{"a1", "b"}
	out <- []string{"a2", "b"}
	out <- []string{"a3", "b"}
	out <- []string{"a4", "b"}
	out <- []string{"a5", "b"}*/
	rand.Seed(int64(time.Now().Nanosecond()))
	to := rand.Intn(30)
	for i := 0; i < to; i++ {
		out <- []string{"a" + fmt.Sprint(i), "b"}
	}
}

func processData(in chan []string, out chan []string, buf chan struct{}, countChan chan int) {
	defer close(out)

	wg := new(sync.WaitGroup)

	for d := range in {
		wg.Add(1)
		buf <- struct{}{}
		go func(data []string) {
			defer wg.Done()
			out <- append(data, "c")
			countChan <- 1
			<-buf
		}(d)
	}
	wg.Wait()
}

/*func printData(in chan []string, countChan chan int)  {
	counter := 0

	//defer func() {
	//	fmt.Println("Total: ", counter)
	//}()

	for {
		select {
		case c, ok := <- countChan:
			if !ok {
				fmt.Println("Total: ", counter)
				return
			}

			counter += c
		case data, ok := <-in:
			if !ok {
				fmt.Println("Total: ", counter)
				return
			}
			fmt.Println(data)
		}
	}
	fmt.Println("Total: ", counter)
}*/

func printData(in chan []string, countChan chan int) {
	counter := 0

	Ololo:
		for {
			select {
			case c, ok := <-countChan:
				if !ok {
					break Ololo
				}
				counter += c
			case data, ok := <-in:
				if !ok {
					break Ololo
				}
				fmt.Println(data)
			}
		}
	fmt.Println("Total: ", counter)
}




/*
func printData(in chan []string) {
	for data := range in {
		fmt.Println(data)
	}
}*/
