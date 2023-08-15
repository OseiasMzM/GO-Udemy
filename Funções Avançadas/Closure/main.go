package main

import "fmt"

func clousure() func() {
	texto := "Função- Clousure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	// Clousure - são básicamente funções que reverenciam variaveis que estão fora do seu corpo
	text := "Função main"
	fmt.Println(text)

	funcaoNova := clousure()
	funcaoNova()
}
