package main

import (
	"fmt"
	"math"
)

type forma interface { // Assinaturas de metodos
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	alturua float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.alturua * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type circulo struct {
	raio float64
}

func main() {
	fmt.Println("Interface")

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
