package main

import (
	"encoding/json"
	"fmt"
	"structsMethods/foo"
)

type MinhaString string

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	foo.Foo
}

func (u User) GetName() string {
	return u.Name
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	foo := foo.NewFoo("Foo")
	fmt.Println(foo.GetName())

	user := User{ // or := User{1, "Cesar"}} // or := &User{1, "Cesar"}}
		Id:   1,
		Name: "Cesar",
	}
	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	fmt.Println(user)
	fmt.Println(user.Name)
	user.UpdateName("Cesar2")
	fmt.Println(user.GetName())
}
