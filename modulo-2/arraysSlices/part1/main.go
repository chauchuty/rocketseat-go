package main

import (
	"fmt"
)

func main() {
	// Array
	arr := [10]int{}
	fmt.Println(arr)
	// Slice
	slice := []int{}
	fmt.Println(slice)

	// Exemplos
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr1[2] = 19
	slice1 := arr1[0:5]
	slice1[0] = 123
	fmt.Println(arr1, slice1)

	slice2 := []int{1, 2, 3, 4, 5}
	slice2 = append(slice2, 6)
	fmt.Println(slice2)

	arr2 := [5]int{1, 2, 3, 4, 5}
	slice3 := arr2[:] // ou slice3 := arr2[0:5]
	fmt.Println(slice3)
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice4 := []int{1, 2, 3, 4, 5}
	slice5 := slice4[:0]
	fmt.Println(slice5, len(slice5), cap(slice5))

	slice6 := []int{}
	fmt.Println(slice6 == nil) // false

	var slice7 []int
	fmt.Println(slice7 == nil) // true
}
