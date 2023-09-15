package main

import (
	"fmt"
)

func main() {

	canal := escrever("Olá Mundo!")
	fmt.Println(<-canal)
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	// Encapsula uma goroutime e retorna um canal
	go func() {
		for {
			canal <- fmt.Sprintf("Valor: %s", texto)
		}
	}()
	return canal
}
