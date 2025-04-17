package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	arr2 := [3]int{1}
	arr3 := [10]int{5: 400, 6: 500}
	arr4 := [10]string{4: "hello, world"}
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
