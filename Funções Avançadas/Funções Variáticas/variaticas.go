package main

import "fmt"

func soma(numeros ...int) int { // Vai receber de 1 a N números
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) { // Limitação, você não pode ter mais de um parametro variatico por função
	// Obrigatoriamente tem que ser o ultimo
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 4, 56)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo!", 10, 2, 4, 6)
}
