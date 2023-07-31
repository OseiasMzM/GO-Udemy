package main

import "fmt"

// As funções em GO podem ter Parâmetros e Retornos
// Parâmetros:  Dados que colocamos nas funções
// Retorno: 	É oque a função devolve

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println("F: ", txt)
	}
	f("Executando Função")

	f("Cahamando resultadoSoma, resultadoSubtracao")
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	f("Executando resultadoSoma1")
	resultadoSoma1, _ := calculosMatematicos(10, 15) // O retorno do função dará dois resultados, passando o "underline" ele igonora o segundo retorno
	fmt.Println(resultadoSoma1)
}

func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

// As funções podem ter mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
