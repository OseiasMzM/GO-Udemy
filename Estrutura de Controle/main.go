package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controles")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { // Criando uma variavel com if init, você limita ela ao escopo do "IF"
		fmt.Println("Numero maior que zero")
	} else if numero < -10 {
		fmt.Println("número é menor que -10")

	} else {
		fmt.Println("Entre 0 e -10")

	}

}
