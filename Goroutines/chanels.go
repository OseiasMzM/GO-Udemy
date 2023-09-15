package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	// for {
	// 	mensagem, aberto := <-canal // Recebendo um valor "Esperando"
	// 	if !aberto {
	// 		fmt.Println(aberto)

	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Passando um valor para a varavel CANAL
		time.Sleep(time.Second)
	}
	close(canal)
}
