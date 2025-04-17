package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um número aleatório será gerado entre 1 e 100. Tente adivinhar!")

	x := rand.Int64N(100) + 1
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Printf("Tentativa %d: ", i+1)
		scanner.Scan()
		chuteInt, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)

		if err != nil {
			fmt.Println("Entrada inválida. Tente novamente.")
			i--
			continue
		}

		chutes[i] = chuteInt

		switch {
		case chuteInt < x:
			fmt.Println("Seu chute é menor que o número secreto.")
		case chuteInt > x:
			fmt.Println("Seu chute é maior que o número secreto.")
		case chuteInt == x:
			fmt.Printf(
				"Parabéns! Você acertou o número secreto. %d\n"+
					"Você fez %d tentativas.\n"+
					"Essas foram suas tentativas: %v\n",
				x,
				len(chutes),
				chutes[:i],
			)
			return
		}
	}

	fmt.Printf(
		"Você não conseguiu adivinhar o número secreto. O número era %d.\n"+
			"Você fez %d tentativas.\n"+
			"Essas foram suas tentativas: %v\n",
		x,
		len(chutes),
		chutes,
	)
	fmt.Println(`###
Acabou o Jogo!
###`)
}
