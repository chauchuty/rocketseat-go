package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	slice2 := [4]int{1, 2, 3, 4}
	fmt.Println(slice, cap(slice))
	foo(slice)
	bar(slice)
	fmt.Println(slice)
	baz2(slice2)
	fmt.Println(slice2)
}

func foo(slice []int) {
	// _ = slice[3] // bound check
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
}

func bar(slice []int) { // slice is a pointer
	slice[0] = 123
}

func baz2(slice [4]int) {
	slice[0] = 123
}
