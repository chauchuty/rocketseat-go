package main

import (
	"fmt"
	"math"
)

func main() {
	m := make(map[string]string)
	m2 := map[string]string{
		"Nome":      "Cesar",
		"Sobrenome": "Chauchuty",
	}
	m3 := map[string][]int{
		"Numeros": {1, 2, 3, 4, 5},
	}

	m4 := make(map[string]string)
	m4["Nome"] = "Cesar"
	valor, ok := m4["Nome"]
	valor2, ok2 := m4["foo"]
	delete(m4, "Nome")
	valor3, ok3 := m4["Nome"]

	m5 := map[string]string{
		"Nome":      "Cesar",
		"Sobrenome": "Chauchuty",
	}
	clear(m5)

	var f float64
	f = math.NaN()
	var f2 float64
	f2 = math.NaN()

	fmt.Println(f == f2)

	m6 := map[float64]string{
		f:  "Cesar",
		f2: "Chauchuty",
	}
	fmt.Println(m6)
	valor4, ok4 := m6[f]
	fmt.Println(valor4, ok4)

	m7 := map[string]string{
		"Nome":      "Cesar",
		"Sobrenome": "Chauchuty",
		"Idade":     "32",
		"Cidade":    "São Paulo",
	}

	for key, value := range m7 {
		fmt.Println(key, value)
		if key == "Idade" {
			delete(m7, key)
		}
	}
	fmt.Println(m7)

	fmt.Println(m)
	fmt.Println(m == nil)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4, valor, ok)
	fmt.Println(valor2, ok2)
	fmt.Println(valor3, ok3)
	fmt.Println(m5, len(m5))
}
