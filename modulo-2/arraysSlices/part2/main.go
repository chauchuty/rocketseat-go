package main

import "fmt"

var filmesNoDb = []string{
	"O Poderoso Chefão",
	"O Poderoso Chefão II",
	"O Poderoso Chefão III",
	"Senhor dos Aneis",
	"Senhor dos Aneis II",
	"Senhor dos Aneis III",
	"Interstellar",
	"A Chegada",
}

func main() {
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// var filmes []string
	filmes := make([]string, 0, len(resultsFromApi)) // 8

	for _, id := range resultsFromApi {
		fmt.Println(len(filmes), cap(filmes))
		filmes = append(filmes, filmesNoDb[id-1])
		fmt.Println(len(filmes), cap(filmes))
	}

	fmt.Println(filmes)

	// Matrix
	matrix := [][]int{}
	matrix3d := [][][]int{}
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})

	matrix3d = append(matrix3d, matrix)
	matrix3d = append(matrix3d, matrix)
}
