package main

import (
	"fmt"
	"math"
)

func main() {
	if 1 < 2 {
		fmt.Println("1 é menor que 2")
	}

	if x := math.Sqrt(4); x < 1 {
		fmt.Println("A raiz quadrada de 4 é menor que 10")
	} else if x > 0 {
		fmt.Println("A raiz quadrada de 4 é maior que 0")
	} else {
		fmt.Println("A raiz quadrada de 4 não é maior que 0")
	}

	if err := doError(); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Sem erro")
	}
}

func doError() error {
	return fmt.Errorf("error")
}
