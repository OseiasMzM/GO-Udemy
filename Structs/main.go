package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	longradouro string
	numero      uint8
}

func main() {
	fmt.Println("Structs")

	var usuario_novo usuario

	usuario_novo.nome = "John"
	usuario_novo.idade = 40

	fmt.Println(usuario_novo)

	endereco_novo := endereco{"Rua Zezé", 202}

	usuario_novo2 := usuario{"Any", 100, endereco_novo}
	fmt.Println(usuario_novo2)

	usuario_novo3 := usuario{nome: "Other"}
	fmt.Println(usuario_novo3)

}
