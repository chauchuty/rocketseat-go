package main

import (
	"fmt"
	"strconv"
)

// bool
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64, uintptr
// byte (uint8)
// rune (int32)
// float32, float64
// complex64, complex128
// string

func main() {
	const y = 10
	var b uint8 = 10
	var x int = 10084
	takeBytes(b)
	f := float64(x) // conversão
	fmt.Println(f)
	// s := string(x) // raridade
	s := strconv.FormatInt(int64(x), 10)
	fmt.Println(s) // A
	fmt.Println(y)
	takeInt32(y)
	takeInt32(10)
	takeFloat64(y)
}

func takeBytes(b byte) {
	fmt.Println(b)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeFloat64(f float64) {
	fmt.Println(f)
}
