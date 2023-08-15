package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array e Slices")

	// Array é uma lista de valores
	// Arrya em go com tamanhos fixos e tipos fixos
	var array1 [5]int

	array1[0] = 1

	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	// ========================================

	// Slice flexivel # Não é um array # Como se fosse uma fatia de um array
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 000, 0, 0, 0, 0, 0, 0}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))  // Devolver o tipo de uma variavel
	fmt.Println(reflect.TypeOf(array3)) // Devolver o tipo de uma variavel

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(array2)
	// ========================================

	// Array Internos
	slice3 := make([]float32, 10, 15) // Tipo ; tamano ; capacidade
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
}
