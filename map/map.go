package main

import (
	"reflect"
	"strings"
	"fmt"
)

func main()  {
	arr := []interface{}{"str", int(22), false}

	for _, a:= range arr {
		switch a.(type) {
		case int:
			a = reflect.ValueOf(a).Int() * 2
		case bool:
			a = !reflect.ValueOf(a).Bool()
		case string:
			a = strings.ToUpper(reflect.ValueOf(a).String())
		}
		fmt.Printf("%#v\n", a)
	}
}
