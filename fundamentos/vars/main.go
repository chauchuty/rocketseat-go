package main

import "fmt"

var idade int

// var nome4 := "Lucas" // não pode usar := para declarar variáveis globais

func main() {
	nome3 := "Lucas"
	var nome, sobrenome = "Pedro", "Silva"
	var (
		nome2      = "João"
		sobrenome2 = "Oliveira"
	)
	fmt.Println(nome, sobrenome, idade)
	fmt.Println(nome2, sobrenome2, idade)
	fmt.Println(nome3)
}
