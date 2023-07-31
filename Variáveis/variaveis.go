package main

import "fmt"

func main() {
	var variavel1 string = "Variiável 1"
	variavel2 := "Variável 2" // Inferência de tipo

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "any"
		variavel4 string = "any"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "any", "any"
	fmt.Println(variavel5, variavel6)

	const constantel string = "Constante 1"

	fmt.Println(constantel)

} // =>
