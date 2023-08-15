package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	usuario := map[string]string{
		"nome":      "John",
		"Sobrenome": "Doe",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "John",
			"ultimo":   "Doe",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

}
