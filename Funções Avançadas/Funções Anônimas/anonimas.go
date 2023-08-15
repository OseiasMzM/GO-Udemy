package main

import "fmt"

func main() {
	// posso ou não armazenar o retorno se ela tiver em uma variavel
	retorno := func(texto string) string { // função anônima
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Parâmetro recebido")

	fmt.Println(retorno)

}
