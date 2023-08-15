package main

import "fmt"

func fibonacci(posicao uint) uint { // Não recomendada se tiver números muitos grandes
	// 1 1 2 3 5 8 13
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	// São funções que chamam elas mesmas
	// Dependesse de uma outra execução dela mesma
	fmt.Println("Funções Recursivas")
	posicao := uint(12)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
