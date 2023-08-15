package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // Busca o endereço na memoria
	// *ponteiro - mostra o valor que está no endereço

	fmt.Println(variavel3, *ponteiro) // desreferenciação

}
