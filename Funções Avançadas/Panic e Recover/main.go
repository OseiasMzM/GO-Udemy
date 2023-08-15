package main

import (
	"fmt"
)

func recuperarExecucao() {
	// if init, cria uma variavel que irá armazenar o resultado da função recover, se esse 'r' for diferente de nulo quer dizer quer a execução foi recuperada
	if r := recover(); r != nil {
		fmt.Println("Tentando recuperar")
	}
}

func AlunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6!")
	// Irá interromper o fluxo normal do seu programa e vai parar tudo

}

func main() {

	fmt.Println(AlunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}
