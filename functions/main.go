package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	sayHi()
	fmt.Println(sum(1, 2))
	a, b := swap(1, 2)
	fmt.Println(a, b)
	res, rem := divide(10, 3)
	fmt.Println(res, rem)
	x := somar(2)(1)
	fmt.Println(x)
	foo := func() {
		fmt.Println("Hello from foo!")
	}
	foo()
	result := sumVariatica(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func sayHi() {
	fmt.Println("Hi!")
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func divide(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return
}

func somar(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func sumVariatica(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}
