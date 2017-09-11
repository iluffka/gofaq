package main

type In interface {
	test1()
	test2()
}

type C struct{}
type E struct {
	C
}

func (_ *C) test1() {
	println("*C")
}

func (_ C) test2() {
	println("C")
}

func main() {
	c := C{}
	e := E{}

	testIn(&c)
	testIn(&e)
}

func testIn(i In) {
	i.test1()
	i.test2()
}