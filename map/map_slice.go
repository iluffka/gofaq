package main

import (
	"fmt"
	"reflect"
)

func main()  {
	a := []int{1,2,3,4,5}
	b := a[1:]
	a[4] = 0
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(a, b)

	c := [5]int{1,2,3,4,5}
	d := c[1:]
	fmt.Println(reflect.TypeOf(c).Kind())
	fmt.Println(c, d)
}
