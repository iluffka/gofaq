package main

type I interface {
	test()
}

type A struct{}
type B struct {
	A
}

func (_ A)test() {
	println("Aaaab")
}

func main()  {
	a := A{}
	b := B{}

	test(a)
	test(b)
	test(&a)
	test(&b)
}

func test(i I) {
	i.test()
}