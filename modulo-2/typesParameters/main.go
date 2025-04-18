// package main

// import (
// 	"fmt"

// 	"golang.org/x/exp/constraints"
// )

// type MeuTipo string

// func (MeuTipo) Foo() {}

// func main() {
// 	// var mt MeuTipo =
// 	// foo(mt)
// 	foo(10)
// 	foo("oi")
// 	foo([]int{1, 2, 3})
// 	var mt MeuTipo = ""
// 	foo(mt)

// 	var ms MyStruct[string] = MyStruct[string]{}
// 	ms.foo()
// }

// type MyConstraint interface {
// 	// Foo()
// 	int | ~string | []int | constraints.Ordered
// }

// func foo[T MyConstraint](arg T) {
// 	fmt.Println(arg)
// }

// type MyStruct[T any] struct {
// 	Foo T
// }

// func (MyStruct[T]) foo() {
// 	fmt.Println("oi")
// }

// package main

// func main() {

// }

// type Foo struct {
// 	Name string
// }

// func (Foo) do() {}

// type Bar struct {
// 	Name string
// }

// func (Bar) do() {}

// type MinhaInterface interface {
// 	Bar | Foo
// 	do()
// }

// func foo[T MinhaInterface](arg T) {
// 	arg.do()
// }

package main

import "slices"

func main() {
	slices.Contains([]int{1, 2, 3}, 3)
	Contains([]int{1, 2, 3}, 3)
}

func Contains[T comparable](s []T, cmp T) bool {
	for _, str := range s {
		if str == cmp {
			return true
		}
	}
	return false
}
