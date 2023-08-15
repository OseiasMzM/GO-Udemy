package main

import (
	"fmt"
)

func funcao1() {
	fmt.Println("Função 1")
}
func funcao2() {
	fmt.Println("Função 2")
}

func AlunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada, Resultado será retornado!")

	fmt.Println("Entrnado na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false

}
func main() {
	// Defar -> "Adiar"
	// defer funcao1() // Adiar a execução do código até o último momento possivel
	// funcao2()
	fmt.Println(AlunoEstaAprovado(4, 9))
}
