package main

import (
	"fmt"
)

type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
}

// func (Dog) Sound() string {
// 	return "Au! Au!"
// }

// func (Dog) Interface() {
// 	fmt.Println("dog interface called")
// }

func (d *Dog) Sound() string {
	if d == nil {
		return "nil"
	}
	return "Au! Au!"
}

type Cat struct {
	Name string
}

func (d *Cat) Sound() string {
	return "Miau! Miau!"
}

func takeEmptyInterface(a any) {
	str, ok := a.(string)
	if ok {
		fmt.Println("string", str)
	}
}

func takeAnimal(a Animal) {
	switch t := a.(type) {
	case *Dog:
		fmt.Println("dog", t.Name)
	case *Cat:
		fmt.Println("cat", t.Name)
	}
}

type Pedro struct{}

func (Pedro) String() string { // stringer
	return "esse é um teste"
}

func main() {
	// Exemplo 1
	// dog := Dog{}
	// whatDoesThisAnimalSay(dog)
	// bar.TakeFoo(dog)

	// Exemplo 2
	// var a Animal           // nil == true
	// var dog *Dog           // null == true
	// a = dog                // a == nil: false
	// fmt.Println(a.Sound()) // out: "Au! Au!"

	// Exemplo 3
	// var a Animal
	// fmt.Println(a.Sound())

	// takeEmptyInterface("teste")
	// takeEmptyInterface(1)
	// takeEmptyInterface(true)
	// takeEmptyInterface(1.2)

	// takeAnimal(&Dog{Name: "Rex"})
	// takeAnimal(&Cat{Name: "Felix"})

	p := Pedro{}
	fmt.Println(p)
}

func whatDoesThisAnimalSay(animal Animal) {
	fmt.Println(animal.Sound())
}

func foo(x interface{}) { // any

}
