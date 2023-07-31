package main

import "fmt"

func main() {

	var num int = 100

	fmt.Println(num)

	var num2 uint = 100
	fmt.Println(num2)

	// INT32 = RUNE
	var num3 rune = 123456
	fmt.Println(num3)

	// BYTE = UINT8
	var num4 byte = 123
	fmt.Println(num4)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 12561653.45
	fmt.Println(numeroReal2)

	// STRINGS
	// Sempre aspas duplas para strings
	// aspas siples ele consideraria o CHAR
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'b'
	fmt.Println(char) // número da tabela ascii

	//

	var texto string // Todo tipo de dado tem seu valor zero (valor inicial), para o tipo string, se torna uma string vazia e para números o valor 'zero'
	fmt.Println(texto)
}
