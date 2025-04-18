package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

func (u User) Foo() {
	fmt.Println("Foo")
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("Não foi possível criar o usuário")
	}
	return &User{}, nil
}

func main() {
	// a := 10
	// b := 0
	// res, err := dividir(a, b)
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }

	// fmt.Println(res)

	user, error := NewUser(true)
	if error != nil {
		println(error.Error())
		return
	}
	user.Foo()

}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Não é possível dividir por zero")
	}

	return a / b, nil
}

// type error interface {
// 	Error() string
// }
