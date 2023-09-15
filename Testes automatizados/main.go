package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoEndereco := enderecos.TipoDeEndereco("Rua dos Bobos")

	fmt.Println(TipoEndereco)
}
