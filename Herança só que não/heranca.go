package main

import "fmt"

type Person struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	Person    // Pegando os campos em pessoa e jogando para estudante
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa := Person{"João", "Pedro", 20, 178}
	fmt.Println(pessoa)

	estudante := estudante{pessoa, "Programação", "YouTube"}
	fmt.Println(estudante)

}
