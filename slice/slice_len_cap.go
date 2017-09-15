package main

import "fmt"

func main() {
	elements := []int{100, 200, 300}

	// Capacity is now 3.
	fmt.Println(cap(elements))

	// Append another element to the slice.
	elements = append(elements, 400)

	// Capacity is now 6—it has doubled.
	fmt.Println(cap(elements))
}

/*func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}*/

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}