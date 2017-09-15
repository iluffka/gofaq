package main

import (
	"fmt"
	"time"
)

/*func main()  {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i) //10, 10, 10 ...
		}()
	}
	time.Sleep(1 * time.Second)
}*/

/*
func main()  {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}*/

/*
func main() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	msg = "ololo"
	time.Sleep(100 * time.Millisecond)
}*/

func main() {
	var msg = "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	msg = "ololo"
	time.Sleep(100 * time.Millisecond)
}