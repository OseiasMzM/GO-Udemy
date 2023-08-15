package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		fmt.Println("Incrementando 'i': ", i)
		i++
		// time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando 'j': ", j)
		// time.Sleep(time.Second)
	}

	nomes := [3]string{"John", "Jesus", "Eu"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)

	}
	for _, nome := range nomes {
		fmt.Println(nome)

	}
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))

	}
	usuario := map[string]string{
		"nome":      "John",
		"sobrenome": "Doe",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)

	}
	for {
		fmt.Println("Loop Infinito...")
	}

}
