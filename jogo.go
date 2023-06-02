package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var plays []int
	var tries, number, randomNumber int
	var retry string
	fmt.Println("Bem-vindo ao jogo da adivinhação!")
	for {
		fmt.Println("Digite um número entre 1 e 100:")
		fmt.Scan(&number)
		randomNumber = rand.Intn(100) + 1
		tries++
		for {
			if number == randomNumber {
				fmt.Println("Parabéns, você acertou!")
				fmt.Printf("Você utilizou %d tentativas.\n", tries)
				plays = append(plays, tries)
				tries = 0
				break
			} else if number > randomNumber {
				fmt.Printf("O número é menor que %d.\n", number)
			} else if number < randomNumber {
				fmt.Printf("O número é maior que %d.\n", number)
			}
			fmt.Println("Digite um número entre 1 e 100:")
			fmt.Scan(&number)
			tries++
		}
		fmt.Println("Você deseja jogar novamente? (s/n):")
		fmt.Scan(&retry)
		if retry != "n" {
			for {
				if retry == "s" {
					break
				}
				fmt.Println("Você deseja jogar novamente? (s/n):")
				fmt.Scan(&retry)
				if retry == "n" {
					break
				}
			}
		}
		if retry == "n" {
			for i := 0; i < len(plays); i++ {
				fmt.Printf("Você utilizou %d tentativas na jogada %d.\n", plays[i], i+1)
			}
			break
		}

	}
}
