package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo\n"
	canal <- "Olá Mundo 2\n"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
