package main

import "fmt"

func add(s []string)  {
	s = append(s, "d")
}

func main()  {
	s := []string{"a", "b", "c"}
	add(s[1:2])
	fmt.Println(s)
}