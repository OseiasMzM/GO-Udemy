package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Segunda"
	case 2:
		return "Terça"
	case 3:
		return "Quarta"
	case 4:
		return "Quinta"
	case 5:
		return "Sexta"
	case 6:
		return "Sabado"
	case 7:
		return "Domingo"
	default:
		return "Número inválido"
	}
}
func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Segunda"
	case numero == 2:
		diaDaSemana = "Terça"
	case numero == 3:
		diaDaSemana = "Quarta"
	case numero == 4:
		diaDaSemana = "Quinta"
	case numero == 5:
		diaDaSemana = "Sexta"
	case numero == 6:
		diaDaSemana = "Sabado"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(22)
	fmt.Println(dia)

	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)
}
