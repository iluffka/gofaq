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

/*func add(i []int) {
	i = append(i, 1)
}

func main() {
	i := make([]int, 0)
	add(i)
	fmt.Println(i)
}*/
