package main

import (
	"errors"
	"fmt"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string {
	return s.msg
}

var ErrNotFound = errors.New(("not found"))

func foo() error {
	return SqrtError{"Não foi possível encontrar o valor"}
}

func foo2() error {
	err := bar()
	if err != nil {
		// return errors.New("deu erro em foo" + err.Error())
		return fmt.Errorf("deu erro em foo: %w", err)
	}
	return nil
}

var ErrQualquer = errors.New("qualquer erro")

func bar() error {
	return ErrQualquer
}

// func raizQuadrada(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, SqrtError{"Não é possível calcular a raiz quadrada de um número negativo"}
// 	}
// 	return math.Sqrt(x), nil
// }

var ErrQualquer2 = errors.New("qualquer erro 2")

func a() error {
	return ErrQualquer
}

func b() error {
	return ErrQualquer2
}

func foo3() error {
	var errorResult error
	if err := a(); err != nil {
		errorResult = errors.Join(err, errorResult)
	}

	if err := b(); err != nil {
		errorResult = errors.Join(err, errorResult)
	}

	return errorResult
}

func main() {
	// x := -1
	// res, err := raizQuadrada(float64(x))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(res)

	// err := foo()

	// var sqrtError *SqrtError

	// if err != nil && errors.As(err, &sqrtError) {
	// 	fmt.Println(sqrtError.msg)
	// 	return
	// }

	// if err != nil && errors.Is(err, ErrNotFound) {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("foi pra fora")

	// err := foo2()
	// if err != nil && errors.Is(err, ErrQualquer) {
	// 	fmt.Println("deu erro: ", err)
	// 	return
	// }

	err := foo3()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrQualquer))
	fmt.Println(errors.Is(err, ErrQualquer2))
}
